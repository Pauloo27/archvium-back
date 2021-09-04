package utils

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func ParseFilePath(str string) (valid bool, username, fullPath string) {
	return parsePath(str, true)
}

func ParseFolderPath(str string) (valid bool, username, fullPath string) {
	return parsePath(str, false)
}

func parsePath(str string, isFile bool) (valid bool, username, fullPath string) {
	// check for a / prefix
	if !strings.HasPrefix(str, "/") {
		return false, "", ""
	}
	// now, remove it to avoid a empty string when splitting
	str = strings.TrimPrefix(str, "/")

	// also remove any / at the end
	str = strings.TrimSuffix(str, "/")

	splittedStr := strings.Split(str, "/")

	for i, folder := range splittedStr {
		if !IsWord(folder) {
			if !isFile || len(splittedStr)-1 != i || !IsValidFileName(folder) {
				return false, "", ""
			}
		}

		if i == 0 {
			username = folder
		}
		fullPath += "/" + folder
	}
	// i've forgot to change valid to true...
	// and spend 30 minutes trying to understrand why it wasnt working...
	/* look at it */
	valid = true // that one line. it could've saved me 30 minutes...
	// naming the variables in the function signature is great for readability
	// but if YOU (aka ME cuz nobody's going to read this)
	// forget about it... it makes you question your life choises
	return
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

var wordRegex = regexp.MustCompile(`^\w+$`)

func IsWord(str string) bool {
	return wordRegex.MatchString(str)
}

var fileNameRegex = regexp.MustCompile(`^[\w|\.]+$`)

func IsValidFileName(str string) bool {
	return fileNameRegex.MatchString(str)
}

func IsNotUnique(err error) bool {
	// FIXME
	// There's no gorm.ErrUnique... so... raw string check?
	return strings.HasPrefix(err.Error(), "ERROR: duplicate key value")
}

var duplicatedKeyRegex = regexp.MustCompile(`"\w+_(\w+)_key"`)

func GetDuplicatedKey(err error) string {
	if !IsNotUnique(err) {
		return ""
	}
	match := duplicatedKeyRegex.FindStringSubmatch(err.Error())

	return match[1]
}

func Validate(a interface{}) *[]*ValidationError {
	var errs []*ValidationError

	validate := validator.New()
	rawErrs := validate.Struct(a)

	if rawErrs == nil {
		return nil
	}

	for _, err := range rawErrs.(validator.ValidationErrors) {
		tag := err.Tag()
		param := err.Param()
		if param != "" {
			errs = append(errs, &ValidationError{
				Field: err.StructField(), Error: Fmt("%s: %s", tag, param),
			})
		} else {
			errs = append(errs, &ValidationError{
				Field: err.StructField(), Error: tag,
			})
		}
	}

	return &errs
}

func ParseAndValidate(payload interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(payload); err != nil {
			return AsError(ctx, http.StatusBadRequest, "Missing payload")
		}

		errs := Validate(payload)
		if errs != nil {
			ctx.Response().SetStatusCode(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{"errors": errs})
		}

		ctx.Locals("payload", payload)
		return ctx.Next()
	}
}
