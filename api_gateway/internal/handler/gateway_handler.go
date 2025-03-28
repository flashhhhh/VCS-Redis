package handler

import (
	"encoding/json"
	"net/http"
	"project/internal/client"
	"strconv"
)

type GatewayHandler struct {
	UserClient client.UserClient
}

func NewGatewayHandler(userClient client.UserClient) *GatewayHandler {
	return &GatewayHandler{
		UserClient: userClient,
	}
}

func (h *GatewayHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	resp, err := h.UserClient.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.Token))
}

func (h *GatewayHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *GatewayHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
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