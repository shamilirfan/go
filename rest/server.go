package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

func Start(cnf config.Config) {
	// port number
	var PORT string = ":" + fmt.Sprintf("%d", config.GetConfig().HttpPort)

	// use middleware
	manager := middleware.NewManager()
	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Logger,
	)

	// create router
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	Routes(mux, manager)

	// server listening
	fmt.Println("Server running on http://localhost" + PORT)
	err := http.ListenAndServe(PORT, wrappedMux)

	// error handling
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
