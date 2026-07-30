package handler

import (
	"net/http"
	"strconv"

	"go-app/internal/service"
)

type GachaHandler struct {
	svc *service.GachaService
}

type GachaSample struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewGachaHandler(svc *service.GachaService) *GachaHandler {
	return &GachaHandler{svc: svc}
}

func (h *GachaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// パラメータからgachaIdを取得
	gachaID, err := strconv.ParseUint(r.PathValue("gacha_id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid gacha_id")
		return
	}

	// 暫定レスポンス
	resp := []GachaSample{
		{ID: gachaID, Name: "暫定ガチャ"},
	}

	writeJSON(w, http.StatusOK, resp)
}