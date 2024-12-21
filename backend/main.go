package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// ルーターを作成
	r := chi.NewRouter()

	// /で受け付けるハンドラーを登録
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello chi"))
		if err != nil {
			// ログにエラーを出力
			log.Printf("failed to write response: %v", err)
		}
	})

	// ポート3333でサーバーを起動
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		// ログにエラーを出力
		log.Printf("failed to start server: %v", err)
	}
}
