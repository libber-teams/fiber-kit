package controller

type Controller interface {
	Routes() []*Route
}
