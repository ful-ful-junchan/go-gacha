package main

import (
	"context"
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

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// ガチャ情報取得
	gachaSvc := service.NewGachaService(db)
	gachaHandler := handler.NewGachaHandler(gachaSvc)
	// 抽選
	drawSvc := service.NewDrawService(db)
	drawHandler := handler.NewDrawHandler(drawSvc);

	// ルーティング定義
	mux.HandleFunc("GET /gachas/{gacha_id}", gachaHandler.ServeHTTP)
	mux.HandleFunc("POST /gachas/{gacha_id}/draw", drawHandler.ServeHTTP)

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