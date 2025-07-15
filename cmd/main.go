package main

import (
	"fmt"
	"net/http"

	"github.com/CreepySunny/rest-api/internal/routes"
)

func main() {
	router := http.NewServeMux()

	routes.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	fmt.Println("Server stopped")
}
