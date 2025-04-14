package handler

import (
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserClient client.UserClient
}

func NewUserHandler(userClient client.UserClient) *UserHandler {
	return &UserHandler{
		UserClient: userClient,
	}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	
	_, err := h.UserClient.CreateUser(username, password, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	resp, err := h.UserClient.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Token:", resp.Token)

	// Write resp.token to bearer token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, resp.Token)))
}

func (h *UserHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	resp, err := h.UserClient.GetUserByID(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"id":       resp.Id,
		"username": resp.Username,
		"name":     resp.Name,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := h.UserClient.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make([]map[string]interface{}, len(resp.Users))
	for i, user := range resp.Users {
		data[i] = map[string]interface{}{
			"id":       user.Id,
			"username": user.Username,
			"name":     user.Name,
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}