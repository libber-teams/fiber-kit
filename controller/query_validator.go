package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/libber-teams/fiber-kit/errors"
)

// ValidateQuery valida a query da requisição e realiza o parse para a struct
func ValidateQueryMiddleware(factory func() interface{}) func(c *fiber.Ctx) error {
	v, ut := getValidator()
	return func(c *fiber.Ctx) error {
		s := factory()
		c.QueryParser(s)
		err := v.Struct(s)
		if err == nil {
			c.Locals(LOCAL_QUERY, s)
			return c.Next()
		}

		e := err.(validator.ValidationErrors)

		validationError := errors.NewValidationError(
			c.Path(),
			status,
		)
		validationError.Map("pt-br", e, ut)
		return c.Status(status).
			JSON(validationError)
	}
}
