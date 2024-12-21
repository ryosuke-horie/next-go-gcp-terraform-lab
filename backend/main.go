package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// ルーターを作成
	r := chi.NewRouter()

	// /で受け付けるハンドラーを登録
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello chi"))
	})

	// ポート3333でサーバーを起動
	http.ListenAndServe(":3333", r)
}
