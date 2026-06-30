package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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

	meta := []map[string]string{}

	files := form.File["files"]

	file_meta := c.FormValue("files_meta")

	log.Println(file_meta)

	err = json.Unmarshal([]byte(file_meta), &meta)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to unmarshal files_meta"})

	}

	if len(meta) != len(files) {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "file meta and file numbers dosen't match"})

		}
	}

	for i, file := range files {
		path := fmt.Sprintf("%s%s/%s", h.config.BaseDir, meta[i]["filepath"], file.Filename)
		obj := models.ObjectModel{
			ObjectName: file.Filename,
			Type:       meta[i]["fileType"],
			BlurHash:   meta[i]["blurHash"],
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			AbsPath:    path,
			OtherInfo:  map[string]any{},
		}

		meta_file := fmt.Sprintf("%s%s/.meta.json", h.config.BaseDir, meta[i]["filepath"])
		b, err := os.ReadFile(meta_file)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to read meta json file"})
		}
		log.Println(string(b))
		var metaJson models.MetaFile

		err = json.Unmarshal(b, &metaJson)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to unmarshal meta file"})
		}

		metaJson.Data = append(metaJson.Data, obj)

		data, _ := json.MarshalIndent(metaJson, "", "  ")

		os.WriteFile(meta_file, data, 0644)

		//	pathToSave := fmt.Sprintf("%s%s", h.config.BaseDir, meta[i]["filepath"])

		err = c.SaveFile(file, path)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to reading meta file"})
		}
	}

	return nil
}
