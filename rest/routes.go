package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func Routes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle("GET /products/{productID}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)
}
