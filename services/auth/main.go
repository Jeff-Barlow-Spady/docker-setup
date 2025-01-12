package main

import (
	"log"
	"os"
	"time"

	"github.com/Jeff-Barlow-Spady/docker-setup/services/auth/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	authService := internal.NewAuthService()

	app.Post("/register", func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		if success := authService.CreateUser(req.Username, req.Password); !success {
			return fiber.NewError(fiber.StatusConflict, "Username already exists")
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User created",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		if !authService.VerifyUser(req.Username, req.Password) {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
		}

		token := authService.CreateToken(req.Username)
		return c.JSON(fiber.Map{
			"status": "success",
			"token":  token,
		})
	})

	app.Get("/verify", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing authorization header")
		}

		token := authHeader[7:] // Remove "Bearer " prefix
		username, valid := authService.VerifyToken(token)
		if !valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		return c.JSON(fiber.Map{
			"status":   "success",
			"username": username,
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Starting auth service on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
