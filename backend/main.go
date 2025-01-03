package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/models"
	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/repository"
)

func main() {
	// ロガーの設定
	models.SetLogger(log.Printf)
	models.SetErrorLogger(log.Printf)

	// DB接続
	db, err := initDB()
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	// リポジトリの初期化
	// タスク
	repo := repository.NewTaskRepository(db)

	// ルーターを作成
	r := NewRouter(repo)

	// PORT環境変数を取得、設定されていない場合は8080を使用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT not set, defaulting to %s", port)
	}

	// すべてのインターフェースで指定されたポートでサーバーを起動
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Listening on %s", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		// ログにエラーを出力
		log.Printf("failed to start server: %v", err)
	}
}
