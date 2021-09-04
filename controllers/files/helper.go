package files

import (
	"net/http"
	"os"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func GetFileByPath(c *fiber.Ctx) (string, error) {
	path := utils.GetTargetPath(c)
	valid, username, fullPath := utils.ParseFilePath(path)
	if !valid {
		return "", utils.AsError(c, http.StatusBadRequest, "Invalid file path")
	}

	if username != c.Locals("user_name") {
		return "", utils.AsError(c, http.StatusForbidden, "You can't use that file")
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))

	stat, err := os.Stat(basePath + path)

	if os.IsNotExist(err) {
		return "", utils.AsError(c, http.StatusNotFound, "File not found")
	}

	if stat.IsDir() {
		return "", utils.AsError(c, http.StatusBadRequest, "The target file is actually a folder")
	}

	return fullPath, nil
}
