package helpers

import (
	"errors"
	"fmt"
	"os"

	"github.com/Adityadangi14/photo_app/config"
)

type Helpers struct {
	conf *config.Config
}

func NewHelper(c *config.Config) Helpers {
	return Helpers{conf: c}
}
func (h *Helpers) DirExists(path string) bool {
	p := fmt.Sprintf("%s%s", h.conf.BaseDir, path)
	info, err := os.Stat(p)

	if err == nil {
		return info.IsDir()
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func (h *Helpers) CreateMetaFile(path string) error {
	content := `
		{
			"prev":false,
			"data":[]
		}
	`

	filePath := fmt.Sprintf("%s/.meta.json", path)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
