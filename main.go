package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Pauloo27/archvium/logger"
	"github.com/Pauloo27/archvium/router"
	"github.com/Pauloo27/archvium/services/db"
	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

var port string

func init() {
	err := godotenv.Load()
	logger.HandleFatal(err, ".env not found, copy the .env.default one")

	port = utils.EnvString("PORT")
	logger.HandleFatal(err, "Web server port is invalid")
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: int(utils.EnvInt64("MAX_FILE_SIZE")),
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: utils.EnvString("FRONTEND_URL"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utils.EnvString("AUTH_JWT_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Next()
		},
	}))

	// Or extend your config for customization
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return !strings.HasPrefix(c.Path(), "/v1/auth/login")
		},
		Max:        10,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			if utils.EnvBool("IS_PROXIED") {
				fmt.Println(c.Get("x-forwarded-for"))
				return c.Get("x-forwarded-for")
			} else {
				return c.IP()
			}
		},
	}))

	router.RouteFor(app)

	logger.HandleFatal(db.Connect(
		utils.EnvString("DB_HOST"),
		utils.EnvString("DB_USER"),
		utils.EnvString("DB_PASSWORD"),
		"archvium",
		utils.EnvString("DB_PORT"),
	), "Cannot connect to db")

	db.Setup()

	app.Listen(utils.Fmt(":%s", port))
}
