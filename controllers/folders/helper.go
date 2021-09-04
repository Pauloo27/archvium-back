package folders

import (
	"net/http"
	"os"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func GetFolderByPath(c *fiber.Ctx, ensureExist bool) (string, error) {
	valid, username, fullPath := utils.ParseFolderPath(utils.GetTargetPath(c))
	if !valid {
		return "", utils.AsError(c, http.StatusBadRequest, "Invalid path")
	}

	if username != c.Locals("user_name") {
		return "", utils.AsError(c, http.StatusForbidden, "You can't see that folder")
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))

	stat, err := os.Stat(basePath + fullPath)

	if ensureExist && os.IsNotExist(err) {
		return "", utils.AsError(c, http.StatusNotFound, "Folder not found")
	}

	if ensureExist && !stat.IsDir() {
		return "", utils.AsError(c, http.StatusBadRequest, "The target folder is actually a file")
	}

	return fullPath, nil
}
