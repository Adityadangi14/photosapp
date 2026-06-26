package main

import (
	"log"

	"github.com/Adityadangi14/photo_app/config"
	"github.com/Adityadangi14/photo_app/internal/server"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config")
	}
	s := server.NewServer(app, &conf)

	s.Run()

}

//2725dd00-626e-43dd-befe-1386a7e81aac
//
//
// mp api(dir path) ->
//
