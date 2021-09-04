package auth

import (
	"net/http"
	"strings"

	"github.com/Pauloo27/archvium/model"
	"github.com/Pauloo27/archvium/services/db"
	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
)

type RegisterPayload struct {
	Username string `validate:"required,min=5,max=32,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=5,max=32"`
}

func RegisterPayloadNormalizer(pl interface{}) {
	payload := pl.(*RegisterPayload)
	payload.Email = strings.ToLower(payload.Email)
	payload.Username = strings.ToLower(payload.Username)
}

func Register(ctx *fiber.Ctx) error {
	if !(ctx.Locals("ENV_AUTH_SELF_REGISTER").(bool)) {
		return utils.AsError(ctx, http.StatusUnauthorized, "Self register is disabled")
	}

	payload := ctx.Locals("payload").(*RegisterPayload)
	user := model.User{
		Email: payload.Email, Password: payload.Password, Username: payload.Username,
	}

	err := db.Connection.Create(&user).Error
	if err != nil {
		if utils.IsNotUnique(err) {
			return utils.AsError(ctx, http.StatusConflict, utils.Fmt(
				"%s already in use", utils.GetDuplicatedKey(err),
			))
		} else {
			return utils.AsError(ctx, http.StatusInternalServerError, err.Error())
		}
	}

	// TODO: create user home folder

	return utils.AsJSON(ctx, http.StatusCreated, user.ToDto())
}
