package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Fmt(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}

func GetTargetPath(c *fiber.Ctx) string {
	return strings.TrimPrefix(c.Path(), strings.TrimSuffix(c.Route().Path, "/*"))
}

func WithoutSlashSuffix(str string) string {
	return strings.TrimSuffix(str, "/")
}

func WithoutSlashPrefix(str string) string {
	return strings.TrimPrefix(str, "/")
}

func GetFileInfo(basePath, path string) (*fiber.Map, error) {
	basePath = WithoutSlashSuffix(basePath)
	path = WithoutSlashPrefix(path)
	stat, err := os.Stat(basePath+"/"+path)
	if err != nil {
		return nil, err
	}

	return &fiber.Map{
		"name":       stat.Name(),
		"isDir":      stat.IsDir(),
		"modifiedAt": stat.ModTime(),
		"size":       stat.Size(),
		"path":       "/"+path,
	}, nil
}
