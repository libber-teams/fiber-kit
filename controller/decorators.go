package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	LOCAL_BODY   = "body"
	LOCAL_QUERY  = "query"
	LOCAL_HEADER = "header"
)

func GetBody[T any](c *fiber.Ctx) T {
	return c.Locals(LOCAL_BODY).(T)
}

func GetQuery[T any](c *fiber.Ctx) T {
	return c.Locals(LOCAL_QUERY).(T)
}

func GetHeaders[T any](c *fiber.Ctx) T {
	return c.Locals(LOCAL_HEADER).(T)
}

// Não válida se o parâmetro é um inteiro.
// Utilize o middleware IntParamMiddleware para validar se o parâmetro é um inteiro.
func GetIntParam(c *fiber.Ctx, param string) int {
	id := c.Params(param)
	i, _ := strconv.Atoi(id)
	return i
}
