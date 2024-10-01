package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAuthRoutes(conn *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", GetLoginHandler(conn))

	return r
}
