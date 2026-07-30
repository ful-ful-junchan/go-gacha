package main

import (
	"log"
	"net/http"

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
	}

	// サーバーを起動
	log.Fatal(srv.ListenAndServe())
}