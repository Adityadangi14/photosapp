package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) CreateDir(c fiber.Ctx) error {

	type Body struct {
		Path string `json:"path"`
	}

	var body Body

	bdy := c.Body()

	err := json.Unmarshal(bdy, &body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "valid path is required",
		})
	}

	if body.Path == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "valid path is required",
		})
	}

	if h.Helper.DirExists(body.Path) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "path already exist",
		})
	}
	fullPath := fmt.Sprintf("%s%s", h.config.BaseDir, body.Path)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create directory",
		})
	}

	err = h.Helper.CreateMetaFile(fullPath)

	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create meta file",
		})
	}
	return c.JSON(fiber.Map{
		"message": "directory created",
		"path":    body.Path,
	})
}
