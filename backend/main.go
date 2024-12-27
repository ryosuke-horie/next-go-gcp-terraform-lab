package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/models"
	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/repository"
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

	// ポート3333でサーバーを起動
	err = http.ListenAndServe(":3333", r)
	if err != nil {
		// ログにエラーを出力
		log.Printf("failed to start server: %v", err)
	}
}
