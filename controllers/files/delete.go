package files

import (
	"net/http"
	"os"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	path, err := GetFileByPath(c)
	if path == "" {
		return err
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))

	err = os.Remove(utils.Fmt("%s/%s", basePath, path))
	if err != nil {
		return utils.AsError(c, http.StatusInternalServerError, "Something went wrong while deleting file from disk")
	}

	return c.SendStatus(http.StatusNoContent)
}
