package service

import (
	"context"
	"database/sql"

	"go-app/internal/repository"
)

type GachaService struct {
	db *sql.DB
}

func NewGachaService(db *sql.DB) *GachaService {
	return &GachaService{db: db}
}

type GachaInfo struct {
	GachaID       uint64
	Name          string
	PityThreshold uint64
}

// GetGachaInfo はガチャの基本情報と、gacha_itemsのweightから導出した提供割合を返す。
func (s *GachaService) GetGachaInfo(ctx context.Context, gachaID uint64) (*GachaInfo, error) {
	gacha, err := repository.FindGachaByID(ctx, s.db, gachaID)
	if err != nil {
		return nil, err
	}

	return &GachaInfo{
		GachaID:       gacha.ID,
		Name:          gacha.Name,
		PityThreshold: gacha.PityThreshold,
	}, nil
}