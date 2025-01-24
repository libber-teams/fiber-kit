package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/libber-teams/fiber-kit/errors"
)

// ValidateBody valida o corpo da requisição e realiza o parse para a struct
func ValidateBodyMiddleware(factory func() interface{}) func(c *fiber.Ctx) error {
	v, t := getValidator()
	return func(c *fiber.Ctx) error {
		s := factory()
		c.BodyParser(s)
		err := v.Struct(s)
		if err == nil {
			c.Locals(LOCAL_BODY, s)
			return c.Next()
		}

		e := err.(validator.ValidationErrors)

		validationError := errors.NewValidationError(
			c.Path(),
			status,
		)
		validationError.Map("pt-br", e, t)
		return c.Status(status).
			JSON(validationError)
	}
}
