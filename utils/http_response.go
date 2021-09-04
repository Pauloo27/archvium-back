package utils

import "github.com/gofiber/fiber/v2"

func AsJSON(c *fiber.Ctx, status int, json fiber.Map) error {
	c.Response().SetStatusCode(status)
	return c.JSON(json)
}

func AsError(c *fiber.Ctx, status int, msg string) error {
	return AsJSON(c, status, fiber.Map{"error": msg})
}

func AsMsg(c *fiber.Ctx, status int, msg string) error {
	return AsJSON(c, status, fiber.Map{"msg": msg})
}
