package main

import (
	"net/http"

	authHttp "github.com/PUCP-Quantum-Optics-Lab/aether-backend/internal/auth/http"
	"github.com/PUCP-Quantum-Optics-Lab/aether-backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	conn := database.GetConnection()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/auth", authHttp.GetAuthRoutes(conn))

	http.ListenAndServe(":3000", r)
}
