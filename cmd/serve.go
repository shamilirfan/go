package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	config := config.GetConfig()
	dbCon, err := db.NewConnection(config.DB)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddlewares(config)

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := product.NewHandler(*middlewares, productRepo)
	userHandler := user.NewHandler(*config, userRepo)

	server := rest.NewServer(
		config,
		productHandler,
		userHandler,
	)

	server.Start()
}
