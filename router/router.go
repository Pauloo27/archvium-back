package router

import (
	authController "github.com/Pauloo27/archvium/controllers/auth"
	filesController "github.com/Pauloo27/archvium/controllers/files"
	usersController "github.com/Pauloo27/archvium/controllers/users"
	foldersController "github.com/Pauloo27/archvium/controllers/folders"
	"github.com/gofiber/fiber/v2"
)

func RouteFor(app *fiber.App) {

	v1 := app.Group("/v1")
	auth := v1.Group("/auth")
	users := v1.Group("/users")
	files := v1.Group("/files")
	folders := v1.Group("/folders")

	auth.Post("/register",
		requireGuest,
		withEnvBool("AUTH_SELF_REGISTER"),
		withPayload(new(authController.RegisterPayload)),
		withPayloadNormalizer(authController.RegisterPayloadNormalizer),
		authController.Register,
	)
	auth.Post("/login",
		requireGuest,
		withPayload(new(authController.LoginPayload)),
		withPayloadNormalizer(authController.LoginPayloadNormalizer),
		authController.Login,
	)

	files.Post("/",
		requireAuth,
		withEnv("STORAGE_ROOT"),
		withEnvInt64("MAX_FILE_SIZE"),
		filesController.Upload,
	)
	files.Get("/info/*",
		requireAuth,
		withEnv("STORAGE_ROOT"),
		filesController.Info,
	)
	files.Get("/download/*",
		requireAuth,
		withEnv("STORAGE_ROOT"),
		filesController.Download,
	)
	files.Delete("/*", 
		requireAuth,
		withEnv("STORAGE_ROOT"),
		filesController.Delete,
	)

	folders.Get("/index/*",
		requireAuth,
		withEnv("STORAGE_ROOT"),
		foldersController.Index,
	)
	folders.Post("/*",
		requireAuth,
		withEnv("STORAGE_ROOT"),
		foldersController.Create,
	)

	users.Get("/@me", requireAuth, usersController.GetMe)
}
