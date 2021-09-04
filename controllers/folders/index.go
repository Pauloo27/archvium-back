package folders

import (
	"net/http"
	"os"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	path, err := GetFolderByPath(c, true)
	if path == "" {
		return err
	}

	basePath := utils.WithoutSlashSuffix(c.Locals("ENV_STORAGE_ROOT").(string))
	files, err := os.ReadDir(basePath + path)
	if err != nil {
		return utils.AsError(c, http.StatusInternalServerError, "Cannot list files in folder")
	}

	var filesInfo = []*fiber.Map{}

	for _, file := range files {
		info, err := utils.GetFileInfo(basePath, path + "/" + file.Name())
		if err != nil {
			return utils.AsError(c, http.StatusInternalServerError, "Cannot get file info")
		}
		filesInfo = append(filesInfo, info)
	}

	return c.JSON(filesInfo)
}
