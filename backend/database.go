package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

// DB初期化処理
func initDB() (*sql.DB, error) {
	// DB接続情報
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "xodb")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")

	var connStr string
	if strings.HasPrefix(dbHost, "/") {
		// Unixソケットを使用する場合
		connStr = fmt.Sprintf("postgres://%s:%s@/%s?host=%s&sslmode=disable", dbUser, dbPassword, dbName, dbHost)
	} else {
		// TCP接続を使用する場合
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	// DB接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("データベースへ接続に失敗しました: %v", err)
		return nil, err
	}

	// データベース接続の確認
	if err := db.Ping(); err != nil {
		log.Fatalf("データベースへのpingに失敗しました: %v", err)
		return nil, err
	}

	return db, nil
}

// 環境変数を取得するヘルパー関数
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
