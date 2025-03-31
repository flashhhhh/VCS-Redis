package main

import (
	"net/http"
	"project/internal/client"
	"project/internal/handler"
)

func main() {
	// Initialize the user client
	userClient := client.NewUserClient("user_service:50051")

	// Initialize the gateway handler
	gatewayHandler := handler.NewGatewayHandler(userClient)

	// Start the HTTP server
	http.HandleFunc("/user/login", gatewayHandler.LoginHandler)
	http.HandleFunc("/user/getUserById", gatewayHandler.GetUserByIdHandler)
	http.HandleFunc("/user/getAllUsers", gatewayHandler.GetAllUsersHandler)

	http.ListenAndServe(":8080", nil)
}