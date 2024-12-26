package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/models"
)

func main() {
	// ロガーの設定
	models.SetLogger(log.Printf)
	models.SetErrorLogger(log.Printf)

	// DB接続情報
	dbUser := "postgres"
	dbPassword := "password"
	dbName := "xodb"
	dbHost := "localhost"
	dbPort := "5432"
	// DB接続文字列作成
	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	// DB接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("データベースへ接続に失敗しました: %v", err)
	}
	defer db.Close()

	// データベース接続の確認
	if err := db.Ping(); err != nil {
		log.Fatalf("データベースへのpingに失敗しました: %v", err)
	}

	// ルーターを作成
	r := chi.NewRouter()

	// /taskでタスク整理CRUDを実装
	// Create
	r.Post("/task", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// ポート3333でサーバーを起動
	err = http.ListenAndServe(":3333", r)
	if err != nil {
		// ログにエラーを出力
		log.Printf("failed to start server: %v", err)
	}
}
