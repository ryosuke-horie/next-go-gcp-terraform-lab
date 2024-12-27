package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/handlers"
)

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Create
	r.Post("/task", handlers.CreateTaskHandler(db))
	// READ
	r.Get("/task", handlers.ListTaskHandler(db))
	// DELETE
	r.Delete("/task", handlers.DeleteTaskHandler(db))

	return r
}
