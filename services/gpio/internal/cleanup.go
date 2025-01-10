package internal

import (
    "os"
    "os/signal"
    "syscall"
    "periph.io/x/conn/v3/gpio"
)

func (g *GPIOManager) InitializeCleanup() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        g.cleanup()
        os.Exit(0)
    }()
}

func (g *GPIOManager) cleanup() {
    g.mu.Lock()
    defer g.mu.Unlock()
    for _, pin := range g.pins {
        pin.Out(gpio.Low)
    }
}
