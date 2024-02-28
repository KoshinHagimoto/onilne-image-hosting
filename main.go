package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"online-image/app/handlers"
	"online-image/app/middleware"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	// この関数は.envファイルを読み込んで環境変数にセットします
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

var db *sql.DB

// initDBはデータベース接続を初期化します
func initDB() {
	var err error
	loadEnv() // .envファイルから接続情報を取得

	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")

	// "ユーザー名:パスワード@tcp(localhost:5555)/データベース名"
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME))
	if err != nil {
		log.Fatal(err)
	}

	// データベース接続を確認
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// テーブルが存在しない場合は作成
	createTable := `
	CREATE TABLE IF NOT EXISTS images (
		id INT AUTO_INCREMENT PRIMARY KEY,
		media_type VARCHAR(255) NOT NULL,
		unique_string VARCHAR(255) NOT NULL,
		file_path VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB() // データベース接続を初期化

	r := mux.NewRouter()

	//CORSミドルウェアを適用
	r.Use(middleware.EnableCORS)

	r.HandleFunc("/upload", handlers.UploadHandler(db)).Methods("POST")
	r.HandleFunc("/images", handlers.ListImageHandler(db)).Methods("GET")
	r.HandleFunc("/image/{mediaType}/{uniqueString}", handlers.ViewHandler).Methods("GET")
	r.HandleFunc("/delete/{uniqueString}", handlers.DeleteHandler).Methods("GET")

	// フロントエンドのファイルを提供
	publicDir := http.Dir("./public/")
	fileServer := http.FileServer(publicDir)
	r.PathPrefix("/").Handler(fileServer)

	http.ListenAndServe(":8080", r)
}
