package main

import (
	"database/sql"
	"log"
)

// DB初期化処理
func initDB() (*sql.DB, error) {
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
		return nil, err
	}

	// データベース接続の確認
	if err := db.Ping(); err != nil {
		log.Fatalf("データベースへのpingに失敗しました: %v", err)
		return nil, err
	}

	return db, nil
}
