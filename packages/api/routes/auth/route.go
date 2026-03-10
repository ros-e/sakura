package auth

import (
	"net/http"
	"os"

	"github.com/ros-e/sakura/packages/api/internal/auth"
	"github.com/ros-e/sakura/packages/api/internal/json"
)

type response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type RegisterPayload struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthHandlers() {
	http.HandleFunc("POST /auth/register", func(w http.ResponseWriter, r *http.Request) {
		var payload RegisterPayload
		if err := json.Read(r, &payload); err != nil {
			json.Write(w, http.StatusBadRequest, response{Error: true, Message: "Invalid request"})
			return
		}
		homedir, err := os.UserHomeDir()
		if err != nil {
			json.Write(w, http.StatusInternalServerError, response{Error: true, Message: "Failed to get home directory"})
			return
		}
		if payload.Password != payload.ConfirmPassword {
			json.Write(w, http.StatusBadRequest, response{Error: true, Message: "Passwords do not match"})
			return
		}
		auth.CreateProfile(payload.Username, payload.Password, homedir+".sakura/config.yaml")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
}
