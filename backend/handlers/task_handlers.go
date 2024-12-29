// handlers/task_handlers.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/models"
	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/repository"
)

type TaskHandler struct {
	Repo repository.TaskRepository
}

// コンストラクタ
func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}

// レスポンス用の構造体
// Detailをstringとして直接レスポンスで返したいため実装
type TaskResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Detail      string    `json:"detail"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// タスク作成処理
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
	}

	// リクエストボディをデコード
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "無効なリクエストボディ", http.StatusBadRequest)
		return
	}

	// タスクを作成
	task := &models.Task{
		Title:       input.Title,
		Detail:      sql.NullString{String: input.Detail, Valid: input.Detail != ""},
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	// Insert
	if err := h.Repo.CreateTask(r.Context(), task); err != nil {
		log.Printf("タスクの挿入に失敗しました。:%v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// TaskResponse にマッピング
	response := TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Detail:      task.Detail.String, // sql.NullString から string を取得
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
	}

	// 作成したタスクをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("レスポンスのエンコードに失敗しました。: %v", err)
	}
}

func (h *TaskHandler) ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Repo.ListTasks(r.Context())
	if err != nil {
		log.Printf("タスクの取得に失敗しました。: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// models.TaskからTaskResponseへのマッピング
	var responseTasks []TaskResponse
	for _, task := range tasks {
		responseTask := TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Detail:      task.Detail.String, // sql.NullStringからstringを取得
			IsCompleted: task.IsCompleted,
			CreatedAt:   task.CreatedAt,
		}
		responseTasks = append(responseTasks, responseTask)
	}

	// レスポンスの設定と送信
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseTasks); err != nil {
		log.Printf("レスポンスのエンコードに失敗しました。: %v", err)
		http.Error(w, "レスポンスの生成に失敗しました。", http.StatusInternalServerError)
		return
	}
}

// タスク削除処理
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// idを受け取り該当のタスクを削除する
	var input struct {
		ID int `json:"id"`
	}

	// リクエストボディをデコード
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "無効なリクエストボディ", http.StatusBadRequest)
		return
	}

	// タスクを削除
	if err := h.Repo.DeleteTask(r.Context(), input.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "タスクが見つかりません", http.StatusNotFound)
		}

		log.Printf("タスクの削除に失敗しました。: %v", err)
		http.Error(w, "タスクの削除に失敗しました。", http.StatusInternalServerError)
		return
	}

	// 削除成功時に204 No Contentを返す
	w.WriteHeader(http.StatusNoContent)
}
