package main

import (
	"log"

	"github.com/Adityadangi14/photo_app/internal/routes"
	"github.com/Adityadangi14/photo_app/internal/server"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	routes.RegisterRoutes(app)

	s := server.NewServer(app)

	err := s.Run()

	if err != nil {
		log.Fatal(err)
	}

}

//2725dd00-626e-43dd-befe-1386a7e81aac
//
//
// mp api(dir path) ->
//
