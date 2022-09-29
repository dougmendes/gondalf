package controller

import "net/http"

type Controller interface {
	AddRoutes(Router)
}

type Router interface {
	Get(pattern string, handler Handler, mw ...Middleware)
}
type Handler func(w http.ResponseWriter, r *http.Request) error

type Middleware func(http.HandlerFunc) http.HandlerFunc
