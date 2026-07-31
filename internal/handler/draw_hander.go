package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"go-app/internal/repository"
	"go-app/internal/service"
)

type DrawHandler struct {
	svc *service.DrawService
}

func NewDrawHandler(svc *service.DrawService) *DrawHandler {
	return &DrawHandler{svc: svc}
}

type drawRequest struct {
	UserID *uint64 `json:"user_id"`
	Count  int     `json:"count"`
}

type drawResultItem struct {
	ItemID   uint64 `json:"item_id"`
	ItemName string `json:"item_name"`
	Rarity   string `json:"rarity"`
	IsPity   bool   `json:"is_pity"`
}

type drawResponse struct {
	Results   []drawResultItem `json:"results"`
	PityCount uint64           `json:"pity_count"`
}

func (h *DrawHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gachaID, err := strconv.ParseUint(r.PathValue("gacha_id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid gacha_id")
		return
	}

	// リクエスト内容をdrawRequestへ展開
	var req drawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.UserID == nil {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	// 抽選実行
	result, err := h.svc.Draw(r.Context(), gachaID, *req.UserID, req.Count)
	switch {
	case errors.Is(err, service.ErrInvalidCount):
		writeError(w, http.StatusBadRequest, "invalid count")
		return
	case errors.Is(err, service.ErrGachaClosed):
		writeError(w, http.StatusConflict, "gacha closed")
		return
	case errors.Is(err, repository.ErrNotFound):
		writeError(w, http.StatusNotFound, "gacha not found")
		return
	case err != nil:
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := drawResponse{PityCount: result.PityCount, Results: make([]drawResultItem, 0, len(result.Results))}
	for _, item := range result.Results {
		resp.Results = append(resp.Results, drawResultItem{
			ItemID:   item.ItemID,
			ItemName: item.ItemName,
			Rarity:   string(item.Rarity),
			IsPity:   item.IsPity,
		})
	}
	writeJSON(w, http.StatusOK, resp)
}