package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"os"
	
	_ "github.com/go-sql-driver/mysql"

	"go-app/internal/handler"
	"go-app/internal/service"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// ガチャ情報取得
	gachaSvc := service.NewGachaService(db)
	gachaHandler := handler.NewGachaHandler(gachaSvc)

	mux.HandleFunc("GET /gachas/{gacha_id}", gachaHandler.ServeHTTP)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// サーバーを起動
	log.Fatal(srv.ListenAndServe())
}