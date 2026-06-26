package server

import (
	"log"

	handlers "github.com/Adityadangi14/photo_app/internal/handlers/object"
	"github.com/gofiber/fiber/v3"
)

type Server interface {
	Run() error
}

func NewServer(app *fiber.App) Server {
	return &server{
		app: app,
	}
}

type server struct {
	app *fiber.App
}

func (s *server) Run() error {

	h := handlers.NewHandler()

	RegisterRoutes(s.app, h)

	log.Fatal(s.app.Listen(":3000"))

	return nil
}
