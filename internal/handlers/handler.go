package handlers

import (
	"github.com/Adityadangi14/photo_app/config"
	"github.com/Adityadangi14/photo_app/internal/helpers"
	"github.com/gofiber/fiber/v3"
)

type Handlers interface {
	Upload(fiber.Ctx) error
	CreateDir(fiber.Ctx) error
}

func NewHandler(conf *config.Config, h *helpers.Helpers) Handlers {
	return &Handler{
		config: conf,
		Helper: h,
	}
}

type Handler struct {
	config *config.Config
	Helper *helpers.Helpers
}
