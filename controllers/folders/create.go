package folders

import (
	"net/http"
	"os"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	path, err := GetFolderByPath(c, false)
	if path == "" {
		return err
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))
	realPath := basePath + path
	if _, err := os.Stat(realPath); !os.IsNotExist(err) {
		return utils.AsError(c, http.StatusConflict, "Folder already exist")
	}

	err = os.MkdirAll(basePath+path, 0700)
	if err != nil {
		return utils.AsError(c, http.StatusInternalServerError,
			"Something went wrong while creating the folder")
	}

	return c.SendStatus(http.StatusCreated)
}
