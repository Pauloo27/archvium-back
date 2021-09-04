package users

import (
	"net/http"

	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	return utils.AsJSON(c, http.StatusOK, fiber.Map{
		"username": c.Locals("user_name"),
		"user_id":  c.Locals("user_id"),
	})
}
