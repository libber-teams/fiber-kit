package controller

import "github.com/gofiber/fiber/v2"

type Route struct {
	Method string
	Path   string

	handler     func(c *fiber.Ctx) error
	middlewares []func(c *fiber.Ctx) error

	hasBodyFactory  bool
	hasQueryFactory bool
	hasHandler      bool

	bodyStructFactory  func() interface{}
	queryStructFactory func() interface{}
}

func NewRoute(method, path string, handler func(c *fiber.Ctx) error) *Route {
	r := &Route{
		Method:      method,
		Path:        path,
		handler:     handler,
		middlewares: []func(c *fiber.Ctx) error{},
	}

	if handler != nil {
		r.hasHandler = true
	}

	return r
}

// WithBodyFactory is a function that receives a factory function that returns a struct that will be used to validate the body.
func (r *Route) WithBodyFactory(factory func() interface{}) *Route {
	r.hasBodyFactory = true
	r.bodyStructFactory = factory
	return r
}

// WithQueryFactory is a function that receives a factory function that returns a struct that will be used to validate the query.
func (r *Route) WithQueryFactory(factory func() interface{}) *Route {
	r.hasQueryFactory = true
	r.queryStructFactory = factory
	return r
}

// Handler is a function that hava the logic of the route
func (r *Route) WithHandler(handler func(c *fiber.Ctx) error) *Route {
	r.hasHandler = true
	r.handler = handler
	return r
}

// WithMiddlewares is a function that receives a list of middlewares that will be used in the route. The middlewares will be applied in the order they are passed.
func (r *Route) WithMiddlewares(middlewares ...func(c *fiber.Ctx) error) *Route {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}
