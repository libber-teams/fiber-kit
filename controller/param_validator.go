package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/libber-teams/fiber-kit/errors"
)

func UUIDParamMiddleware(param string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if !isUUIDv4(id) {
			return errors.NewHttpBadRequestError(c.Path()).
				SetType("invalid_param").
				SetTitle(fmt.Sprintf("Parâmetro inválido")).
				SetDetail(fmt.Sprintf("O parâmetro %s deve ser um UUIDv4 válido", param))
		}
		return c.Next()
	}
}

func IntParamMiddleware(param string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if isInteger(id) {
			return errors.NewHttpBadRequestError(c.Path()).
				SetType("invalid_param").
				SetTitle(fmt.Sprintf("Parâmetro inválido")).
				SetDetail(fmt.Sprintf("O parâmetro %s deve ser um número inteiro", param))
		}
		return c.Next()
	}
}

func isInteger(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
