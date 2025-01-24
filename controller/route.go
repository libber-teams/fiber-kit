package controller

import "github.com/gofiber/fiber/v2"

type Route struct {
	Method string
	Path   string

	handler func(c *fiber.Ctx) error

	hasBodyFactory  bool
	hasQueryFactory bool
	hasHandler      bool

	bodyStructFactory  func() interface{}
	queryStructFactory func() interface{}
}

func NewRoute(method, path string, handler func(c *fiber.Ctx) error) *Route {
	r := &Route{
		Method:  method,
		Path:    path,
		handler: handler,
	}

	if handler != nil {
		r.hasHandler = true
	}

	return r
}

func (r *Route) WithBodyFactory(factory func() interface{}) *Route {
	r.hasBodyFactory = true
	r.bodyStructFactory = factory
	return r
}

func (r *Route) WithQueryFactory(factory func() interface{}) *Route {
	r.hasQueryFactory = true
	r.queryStructFactory = factory
	return r
}

func (r *Route) WithHandler(handler func(c *fiber.Ctx) error) *Route {
	r.hasHandler = true
	r.handler = handler
	return r
}
