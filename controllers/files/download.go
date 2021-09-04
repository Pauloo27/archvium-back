package files

import (
	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {
	path, err := GetFileByPath(c)
	if path == "" {
		return err
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))

	return c.SendFile(basePath+path, false)
}
