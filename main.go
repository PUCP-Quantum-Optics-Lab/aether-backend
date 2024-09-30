package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
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

		w.Write([]byte(loginRequest.Password))
	})

	http.ListenAndServe(":3000", r)
}
