package server

import (
	handlers "github.com/Adityadangi14/photo_app/internal/handlers/object"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, h handlers.Handler) {
	r := app.Group("/api/v1")

	r.Post("/upload", h.Upload)

}
