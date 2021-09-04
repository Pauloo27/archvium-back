package files

import (
	"net/http"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {
	path, err := GetFileByPath(c)
	if path == "" {
		return err
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))

	info, err := utils.GetFileInfo(basePath, path)
	if err != nil {
		return utils.AsError(c, http.StatusInternalServerError, "Something went wrong while getting file info")
	}
	return utils.AsJSON(c, http.StatusOK, *info)
}
