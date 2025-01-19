package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	addr string
	db   *sqlx.DB
}

func NewServer(addr string, db *sqlx.DB) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Run() error {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return http.ListenAndServe(s.addr, router)
}
