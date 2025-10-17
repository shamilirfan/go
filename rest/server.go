package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	config         *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	config *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		config:         config,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	// port number
	var PORT string = ":" + fmt.Sprintf("%d", config.GetConfig().HttpPort)

	// use middleware
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.CorsWithPreflight,
		middlewares.Logger,
	)

	// create router
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	// server listening
	fmt.Println("Server running on http://localhost" + PORT)
	err := http.ListenAndServe(PORT, wrappedMux)

	// error handling
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
