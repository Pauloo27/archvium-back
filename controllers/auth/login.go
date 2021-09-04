package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Pauloo27/archvium/model"
	"github.com/Pauloo27/archvium/services/db"
	"github.com/Pauloo27/archvium/utils"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Username string `validate:"required,min=5,max=32"`
	Password string `validate:"required,min=5,max=32"`
}

func LoginPayloadNormalizer(pl interface{}) {
	payload := pl.(*LoginPayload)
	payload.Username = strings.ToLower(payload.Username)
}

func Login(ctx *fiber.Ctx) error {
	payload := ctx.Locals("payload").(*LoginPayload)

	var user model.User
	err := db.Connection.First(&user,
		"username = ? AND password = ?", payload.Username, utils.HashPassword(payload.Password),
	).Error

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return utils.AsError(ctx, http.StatusNotFound, "Invalid password or username")
		} else {
			return utils.AsError(ctx, http.StatusInternalServerError, err.Error())
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// TODO: refresh
	expiresAt := time.Now().Add(time.Hour * 5).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expiresAt
	claims["username"] = user.Username
	claims["id"] = user.ID

	t, err := token.SignedString([]byte(utils.EnvString("AUTH_JWT_SECRET")))

	return utils.AsJSON(ctx, http.StatusCreated, fiber.Map{
		"token":     t,
		"expiresAt": expiresAt,
		"prefix":    "Bearer",
	})
}
