package handlers

import "github.com/gofiber/fiber/v3"

type Handler interface {
	Upload(fiber.Ctx) error
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct {
}

func (h *handler) Upload(c fiber.Ctx) error {
	return nil
}
