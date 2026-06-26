package handlers

import (
	"encoding/json"

	"github.com/Adityadangi14/photo_app/config"
	models "github.com/Adityadangi14/photo_app/internal/model"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	Upload(fiber.Ctx) error
}

func NewHandler(conf *config.Config) Handler {
	return &handler{
		config: conf,
	}
}

type handler struct {
	config *config.Config
}

func (h *handler) Upload(c fiber.Ctx) error {

	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid Multipart Form"})
	}

	meta := []map[string]any{}

	files := form.File["files"]

	file_meta := c.FormValue("files_meta")

	byt, err := json.Marshal(file_meta)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to marshal files_meta"})
	}

	err = json.Unmarshal(byt, &meta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to unmarshal files_meta"})

	}

	var objModel []models.ObjectModel

	for f := range meta {

	}

	// for fh := range files {
	// 	fileName :=
	// }

	// fmt.Println(c.FormValue("files_meta"))

	// dst := fmt.Sprintf("./base_folder/%s-%s", uuid.NewString(), data[0].Filename)

	// c.SaveFile(data[0], dst)

	return nil
}
