package main

import (
	"log"
	"net/http"
	"time"

	"go-app/internal/handler"
	"go-app/internal/service"
)

func main() {
	mux := http.NewServeMux()

	// ガチャ情報取得
	gachaSvc := &service.GachaService{}
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