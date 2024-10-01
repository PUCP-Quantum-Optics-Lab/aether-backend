package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PUCP-Quantum-Optics-Lab/aether-backend/internal/auth/data"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func GetLoginHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var loginRequest LoginRequest
		if err := json.Unmarshal(body, &loginRequest); err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		user, err := data.GetUser(conn, loginRequest.Username)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed querying user: %v\n", err)
			w.WriteHeader(500)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(loginRequest.Password), []byte(user.Password))
		if err == nil {
			fmt.Fprintf(os.Stderr, "Failed password check: %v\n", err)
			return
		}

		type LoginResponse struct {
			Ok bool `json:"ok"`
		}

		response, err := json.Marshal(LoginResponse{Ok: true})

		w.Write(response)
	}
}
