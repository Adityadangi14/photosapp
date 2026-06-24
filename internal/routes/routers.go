package routes

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	r := app.Group("/api/v1")

	r.Post("/upload", func() {})
}
