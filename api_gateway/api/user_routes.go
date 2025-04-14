package api

import (
	"gateway/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, gatewayHandler *handler.UserHandler) {
	r.HandleFunc("/user/create", gatewayHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/login", gatewayHandler.LoginHandler).Methods("POST")
	r.HandleFunc("/user/getUserById", gatewayHandler.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/user/getAllUsers", gatewayHandler.GetAllUsersHandler).Methods("GET")
}