package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/libber-teams/fiber-kit/errors"
)

func UUIDHeaderMiddleware(headerKey string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Get(headerKey)
		if !isUUIDv4(id) {
			return errors.NewHttpBadRequestError(c.Path()).
				SetType("invalid_header").
				SetTitle(fmt.Sprintf("Header inválido")).
				SetDetail(fmt.Sprintf("O header %s deve ser um UUIDv4 válido.", headerKey))
		}
		return c.Next()
	}
}

func IntHeaderMiddleware(headerKey string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Get(headerKey)
		if isInteger(id) {
			return errors.NewHttpBadRequestError(c.Path()).
				SetType("invalid_header").
				SetTitle(fmt.Sprintf("Header inválido")).
				SetDetail(fmt.Sprintf("O header %s deve ser um número inteiro.", headerKey))
		}
		return c.Next()
	}
}

func ExistHeaderMiddleware(headerKey string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Get(headerKey)
		if id == "" {
			return errors.NewHttpBadRequestError(c.Path()).
				SetType("invalid_header").
				SetTitle(fmt.Sprintf("Header inválido")).
				SetDetail(fmt.Sprintf("O header %s é obrigatório.", headerKey))
		}
		return c.Next()
	}
}
