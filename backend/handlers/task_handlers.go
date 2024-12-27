// handlers/task_handlers.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/models"
)

// タスク作成処理
func CreateTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		if err := task.Insert(r.Context(), db); err != nil {
			log.Printf("タスクの挿入に失敗しました。:%v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// 作成したタスクをレスポンスとして返す
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(task); err != nil {
			log.Printf("レスポンスのエンコードに失敗しました。: %v", err)
		}
	}
}

func ListTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT * FROM Tasks;`)
		if err != nil {
			log.Fatalf("DB読み込みに失敗しました。: %v", err)
			http.Error(w, "DB読み込みに失敗しました。", http.StatusInternalServerError)
		}
		defer rows.Close()

		var tasks []models.Task

		// 各行を反復処理
		for rows.Next() {
			var task models.Task
			err := rows.Scan(
				&task.ID,
				&task.Title,
				&task.Detail,
				&task.IsCompleted,
				&task.CreatedAt,
			)
			if err != nil {
				log.Printf("行のスキャンに失敗しました。: %v", err)
				http.Error(w, "データの処理中にエラーが発生しました。", http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}

		// 反復中にエラーが発生していないか確認
		if err = rows.Err(); err != nil {
			log.Printf("rowsのイテレーション中にエラーが発生しました。: %v", err)
			http.Error(w, "データの取得中にエラーが発生しました。", http.StatusInternalServerError)
			return
		}

		// レスポンスの設定と送信
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			log.Printf("レスポンスのエンコードに失敗しました。: %v", err)
			http.Error(w, "レスポンスの生成に失敗しました。", http.StatusInternalServerError)
			return
		}
	}
}

func DeleteTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID int `json:"id"`
		}

		// リクエストボディをデコード
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "無効なリクエストボディ", http.StatusBadRequest)
			return
		}

		// IDを指定したTaskを作成し削除
		task := models.Task{
			ID: input.ID,
		}
		// タスクを削除
		if err := task.Delete(r.Context(), db); err != nil {
			log.Printf("タスクの削除に失敗しました。: %v", err)
			http.Error(w, "タスクの削除に失敗しました。", http.StatusInternalServerError)
			return
		}

		// 削除成功時に204 No Contentを返す
		w.WriteHeader(http.StatusNoContent)
	}
}
