package main

import (
	"github.com/Adityadangi14/photo_app/internal/server"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	s := server.NewServer(app)

	s.Run()

}

//2725dd00-626e-43dd-befe-1386a7e81aac
//
//
// mp api(dir path) ->
//
