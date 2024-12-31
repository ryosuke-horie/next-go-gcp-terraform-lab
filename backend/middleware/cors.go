package middleware

import (
	"net/http"
)

// CORSMiddleware はCORSヘッダーを設定し、OPTIONSリクエストに対応します。
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 必要に応じてオリジンを動的に設定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// OPTIONSリクエストの場合、200を返して終了
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 次のハンドラーに処理を委譲
		next.ServeHTTP(w, r)
	})
}
