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

		var fr fiber.Router
		hasError := false
		switch route.Method {
		case "GET":
			fr = f.Get(route.Path)
		case "POST":
			fr = f.Post(route.Path)
		case "PUT":
			fr = f.Put(route.Path)
		case "DELETE":
			fr = f.Delete(route.Path)
		case "PATCH":
			fr = f.Patch(route.Path)
		case "OPTIONS":
			fr = f.Options(route.Path)
		case "HEAD":
			fr = f.Head(route.Path)
		case "CONNECT":
			fr = f.Connect(route.Path)
		case "TRACE":
			fr = f.Trace(route.Path)
		default:
			errors = append(errors, fmt.Sprintf("Method %s not supported in %s", route.Method, route.Path))
			hasError = true
		}
		if hasError {
			continue
		}

		if route.hasBodyFactory {
			fr.Use(ValidateBodyMiddleware(route.bodyStructFactory))
		}
		if route.hasQueryFactory {
			fr.Use(ValidateQueryMiddleware(route.queryStructFactory))
		}

		fr.Use(route.handler)
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
