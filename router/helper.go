package router

import (
	"github.com/Pauloo27/archvium/utils"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func isAuthed(c *fiber.Ctx) bool {
	return c.Locals("user") != nil
}

func withPayload(payload interface{}) fiber.Handler {
	return utils.ParseAndValidate(payload)
}

type Normalizer func(payload interface{})

func withPayloadNormalizer(normalizer Normalizer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		normalizer(c.Locals("payload"))
		return c.Next()
	}
}

func withEnv(name string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("ENV_"+name, utils.Env(name))
		return c.Next()
	}
}

func withEnvBool(name string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("ENV_"+name, utils.EnvBool(name))
		return c.Next()
	}
}

func withEnvInt64(name string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("ENV_"+name, utils.EnvInt64(name))
		return c.Next()
	}
}

func requireAuth(c *fiber.Ctx) error {
	if isAuthed(c) {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		c.Locals("user_name", claims["username"].(string))
		c.Locals("user_id", int(claims["id"].(float64)))
		return c.Next()
	}
	return utils.AsError(c, fiber.StatusUnauthorized, "Authentication is required")
}

func requireGuest(c *fiber.Ctx) error {
	if !isAuthed(c) {
		return c.Next()
	}
	return utils.AsError(c, fiber.StatusForbidden, "Being unauthenticated is required")
}
