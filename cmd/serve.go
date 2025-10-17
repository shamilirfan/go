package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	config := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(config)

	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(*middlewares, productRepo)

	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(config, productHandler, userHandler)

	server.Start()
}
