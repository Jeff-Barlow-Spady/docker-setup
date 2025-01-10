package internal

import (
    "fmt"
    "sync"
    "log"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "periph.io/x/conn/v3/gpio"
    "periph.io/x/conn/v3/gpio/gpioreg"
    "periph.io/x/host/v3"
)

var (
    gpioOperations = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "gpio_operations_total",
            Help: "Total GPIO operations",
        },
        []string{"operation", "pin"},
    )

    gpioState = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "gpio_pin_state",
            Help: "Current GPIO pin state",
        },
        []string{"pin"},
    )
)

type GPIOManager struct {
    mu         sync.RWMutex
    pins       map[int]gpio.PinIO
    callbacks []func(int, bool)
}

func NewGPIOManager() *GPIOManager {
    state, err := host.Init()
    if err != nil {
        log.Fatalf("Failed to initialize periph: %v", err)
    }
    log.Printf("Loaded drivers: %s", state.Loaded)
    
    manager := &GPIOManager{
        pins:       make(map[int]gpio.PinIO),
        callbacks:  make([]func(int, bool), 0),
    }
    
    manager.InitializeCleanup()
    return manager
}

func (g *GPIOManager) SetupPin(pin int, direction string) error {
    g.mu.Lock()
    defer g.mu.Unlock()

    // Check if pin is already configured
    if _, exists := g.pins[pin]; exists {
        return fmt.Errorf("pin %d is already configured", pin)
    }

    // Get the pin
    p := gpioreg.ByNumber(pin)
    if p == nil {
        return fmt.Errorf("failed to find pin %d", pin)
    }

    // Configure the pin based on direction
    var err error
    switch direction {
    case "out":
        err = p.Out(gpio.Low)
    case "in":
        err = p.In(gpio.Float, gpio.NoEdge)
    default:
        return fmt.Errorf("invalid direction: %s", direction)
    }

    if err != nil {
        return fmt.Errorf("failed to configure pin %d: %v", pin, err)
    }

    g.pins[pin] = p
    gpioOperations.WithLabelValues("setup", fmt.Sprintf("%d", pin)).Inc()
    return nil
}

func (g *GPIOManager) WritePin(pin int, value bool) error {
    g.mu.Lock()
    defer g.mu.Unlock()

    p, exists := g.pins[pin]
    if !exists {
        return fmt.Errorf("pin %d not configured", pin)
    }

    level := gpio.Low
    if value {
        level = gpio.High
    }

    err := p.Out(level)
    if err != nil {
        return err
    }

    gpioOperations.WithLabelValues("write", fmt.Sprintf("%d", pin)).Inc()
    gpioState.WithLabelValues(fmt.Sprintf("%d", pin)).Set(boolToFloat64(value))

    // Notify callbacks
    for _, cb := range g.callbacks {
        cb(pin, value)
    }

    return nil
}

func (g *GPIOManager) ReadPin(pin int) (bool, error) {
    g.mu.RLock()
    defer g.mu.RUnlock()

    p, exists := g.pins[pin]
    if !exists {
        return false, fmt.Errorf("pin %d not configured", pin)
    }

    gpioOperations.WithLabelValues("read", fmt.Sprintf("%d", pin)).Inc()
    return p.Read() == gpio.High, nil
}

func (g *GPIOManager) RegisterCallback(callback func(int, bool)) {
    g.mu.Lock()
    defer g.mu.Unlock()
    g.callbacks = append(g.callbacks, callback)
}

func boolToFloat64(b bool) float64 {
    if b {
        return 1
    }
    return 0
}
