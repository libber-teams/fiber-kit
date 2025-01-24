package errors

import (
	"github.com/gofiber/fiber/v2"
)

type HttpErrorMiddleware struct {
}

func NewHttpErrorMiddleware() *HttpErrorMiddleware {
	return &HttpErrorMiddleware{}
}

func (m *HttpErrorMiddleware) Handle(c *fiber.Ctx) error {
	err := c.Next()

	if err == nil {
		return nil
	}

	var httpError *HttpError

	if h, ok := err.(*HttpError); !ok {
		return nil
	} else {
		httpError = h
	}

	return c.Status(httpError.Status).JSON(httpError)
}
