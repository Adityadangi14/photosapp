package server

import (
	handlers "github.com/Adityadangi14/photo_app/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, h handlers.Handlers) {
	r := app.Group("/api/v1")

	r.Post("/upload", h.Upload)

	action := r.Group("/create")

	action.Post("/dir", h.CreateDir)

}
