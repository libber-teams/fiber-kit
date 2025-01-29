package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MountController(controller Controller, f *fiber.App) error {
	errors := []string{}
	for _, route := range controller.Routes() {
		if !route.hasHandler {
			panic(fmt.Errorf("Route %s does not have a handler", route.Path))
		}

		hasError := false
		switch route.Method {
		case "GET":
			f.Get(route.Path, wrapHandler(route)...)
		case "POST":
			f.Post(route.Path, wrapHandler(route)...)
		case "PUT":
			f.Put(route.Path, wrapHandler(route)...)
		case "DELETE":
			f.Delete(route.Path, wrapHandler(route)...)
		case "PATCH":
			f.Patch(route.Path, wrapHandler(route)...)
		case "OPTIONS":
			f.Options(route.Path, wrapHandler(route)...)
		case "HEAD":
			f.Head(route.Path, wrapHandler(route)...)
		case "CONNECT":
			f.Connect(route.Path, wrapHandler(route)...)
		case "TRACE":
			f.Trace(route.Path, wrapHandler(route)...)
		default:
			errors = append(errors, fmt.Sprintf("Method %s not supported in %s", route.Method, route.Path))
			hasError = true
		}
		if hasError {
			continue
		}
	}

	if len(errors) > 0 {
		msg := ""
		for _, err := range errors {
			msg += err + "\n"
		}

		return fmt.Errorf(msg)
	}

	return nil
}

func wrapHandler(route *Route) []fiber.Handler {
	handlers := []fiber.Handler{}
	if route.hasBodyFactory {
		handlers = append(handlers, ValidateBodyMiddleware(route.bodyStructFactory))
	}
	if route.hasQueryFactory {
		handlers = append(handlers, ValidateQueryMiddleware(route.queryStructFactory))
	}

	handlers = append(handlers, route.handler)
	return handlers
}
