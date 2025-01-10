package main

import (
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/Jeff-Barlow-Spady/docker-setup/services/metrics/internal"
)

func main() {
    app := fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }
            return c.Status(code).JSON(fiber.Map{
                "error": err.Error(),
            })
        },
    })

    app.Use(recover.New())
    app.Use(logger.New())

    collector := internal.NewMetricsCollector()
    
    // Get update interval from environment
    intervalStr := os.Getenv("METRICS_UPDATE_INTERVAL")
    updateInterval := 15 * time.Second
    if intervalStr != "" {
        if parsed, err := time.ParseDuration(intervalStr); err == nil {
            updateInterval = parsed
        }
    }
    
    // Update metrics periodically
    go func() {
        ticker := time.NewTicker(updateInterval)
        defer ticker.Stop()
        
        for range ticker.C {
            if err := collector.UpdateMetrics(); err != nil {
                log.Printf("Error updating metrics: %v", err)
            }
        }
    }()

    app.Get("/metrics", func(c *fiber.Ctx) error {
        metrics := collector.GetMetrics()
        return c.JSON(metrics)
    })

    app.Get("/health", func(c *fiber.Ctx) error {
        health := collector.GetHealth()
        status := fiber.StatusOK
        if health.Status == "degraded" {
            status = fiber.StatusServiceUnavailable
        }
        return c.Status(status).JSON(health)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    log.Printf("Starting metrics service on port %s with update interval %s", port, updateInterval)
    log.Fatal(app.Listen(":" + port))
}
