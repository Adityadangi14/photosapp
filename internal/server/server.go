package server

import (
	"log"

	"github.com/Adityadangi14/photo_app/config"
	handlers "github.com/Adityadangi14/photo_app/internal/handlers"
	"github.com/Adityadangi14/photo_app/internal/helpers"
	"github.com/gofiber/fiber/v3"
)

type Server interface {
	Run() error
}

func NewServer(app *fiber.App, conf *config.Config) Server {
	return &server{
		app:    app,
		config: conf,
	}
}

type server struct {
	app    *fiber.App
	config *config.Config
}

func (s *server) Run() error {

	helpers := helpers.NewHelper(s.config)
	h := handlers.NewHandler(s.config, &helpers)

	RegisterRoutes(s.app, h)

	log.Fatal(s.app.Listen(":3000"))

	return nil
}
