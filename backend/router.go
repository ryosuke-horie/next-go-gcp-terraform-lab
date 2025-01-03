package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/handlers"
	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/middleware"
	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/repository"
)

func NewRouter(repo repository.TaskRepository) http.Handler {
	r := chi.NewRouter()

	// ハンドラの初期化処理
	taskHandler := handlers.NewTaskHandler(repo)
	// CORSミドルウェアを全ルートに適用
	r.Use(middleware.CORSMiddleware)

	// Create
	r.Post("/task", taskHandler.CreateTaskHandler)
	// Read
	r.Get("/task", taskHandler.ListTaskHandler)
	// Delete
	r.Delete("/task", taskHandler.DeleteTaskHandler)
	// Update
	r.Put("/task", taskHandler.UpdateTaskHandler)

	return r
}
