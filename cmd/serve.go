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
	userRepo := repo.NewUserRepo()

	productHandler := product.NewHandler(*middlewares, productRepo)
	userHandler := user.NewHandler(*config, userRepo)

	server := rest.NewServer(
		config,
		productHandler,
		userHandler,
	)

	server.Start()
}
