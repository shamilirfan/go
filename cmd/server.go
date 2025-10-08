package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Server() {
	// port number
	var PORT string = ":8080"

	// call routes/endpoint
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
	}
}
