package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler // Middleware, it is called signature

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{}
}

func (manager *Manager) Use(middlewares ...Middleware) *Manager {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
	return &Manager{}
}

func (manager *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	// for _, middleware := range middlewares {
	// 	h = middleware(h)
	// }

	return h
}

func (manager *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	for _, middleware := range manager.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
