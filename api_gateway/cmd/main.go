package main

import (
	"context"
	"fmt"
	"gateway/api"

	"gateway/internal/client"
	"gateway/internal/handler"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/flashhhhh/pkg/env"
	"github.com/flashhhhh/pkg/jwt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func checkCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("Origin:", origin)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Vary", "Origin")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}

		next.ServeHTTP(w, r)
	})
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}

		token := authHeader[1]
		data, err := jwt.ValidateToken(token)
		if err != nil {
			fmt.Println("Invalid token:", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token"))
			return
		}

		userIdFloat, ok := data["id"].(float64)
		if !ok {
			fmt.Println("userId not found in token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token Data"))
			return
		}

		userId := strconv.Itoa(int(userIdFloat))

		// Pass userId to the next handler using context
		ctx := context.WithValue(r.Context(), "userId", userId)

		// Log userId for debugging
		fmt.Println("Passing userId to next handler:", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	var environment string

	args := os.Args
	if len(args) > 1 {
		environment = args[1]
		println("Running in environment: ", environment)
	} else {
		println("No environment specified, defaulting to local")
		environment = "local"
	}

	// Load environment variables from the specified file
	envFile := "config/" + environment + ".env"
	env.LoadEnv(envFile)
	log.Println("Environment variables loaded from file:", envFile)

	// Set up the user service address and port
	userServiceAddress := env.GetEnv("USER_SERVICE_HOST", "localhost") + ":" + env.GetEnv("USER_SERVICE_PORT", "50051")
	log.Println("User service address:", userServiceAddress)
	serverPort := env.GetEnv("GATEWAY_PORT", "1906")

	// Create userHandler
	userClient := client.NewUserClient(userServiceAddress)
	userHandler := handler.NewUserHandler(userClient)

	// Start the HTTP server
	r := mux.NewRouter()
	api.RegisterRoutes(r, userHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins, change this for security
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	log.Println("Starting server on port", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, checkCORS(corsHandler)))
}