package server

import (
	"log"

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

	log.Fatal(s.app.Listen(":3000"))

	return nil
}
