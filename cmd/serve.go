package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	config := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(config)
	productHandler := product.NewHandler(*middlewares)
	userHandler := user.NewHandler()
	server := rest.NewServer(config, productHandler, userHandler)

	server.Start()
}
