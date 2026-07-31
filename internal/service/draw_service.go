package service

import (
	"time"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-app/internal/model"
	"go-app/internal/repository"
)

type DrawService struct {
	db  *sql.DB
}

func NewDrawService(db *sql.DB) *DrawService {
	return &DrawService{db: db}
}

type DrawResultItem struct {
	ItemID   uint64
	ItemName string
	Rarity   model.Rarity
	IsPity   bool
}

type DrawResult struct {
	Results   []DrawResultItem
	PityCount uint64
}

var ErrInvalidCount = errors.New("invalid draw count")
var ErrGachaClosed = errors.New("gacha closed")
var ErrGachaItemsNone = errors.New("gacha items not found")

// ガチャ抽選処理
func (s *DrawService) Draw(ctx context.Context, gachaID, userID uint64, count int) (*DrawResult, error) {
	// 実行回数チェック
	if count != 1 && count != 10 {
		return nil, ErrInvalidCount
	}

	// トランザクション開始
	tans, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tans.Rollback() }()

	// ガチャ情報取得
	gacha, err := repository.FindGachaByID(ctx, tans, gachaID)
	if err != nil {
		return nil, err
	}

	// ガチャ開催判定
	now := time.Now()
	if now.Before(gacha.StartsAt) || now.After(gacha.EndsAt) {
		return nil, ErrGachaClosed
	}

	// 排出アイテム一覧を取得
	gachaItems, err := repository.FindGachaItemsByGachaID(ctx, tans, gachaID)
	if err != nil {
		return nil, err
	}
	if len(gachaItems) == 0 {
		return nil, ErrGachaItemsNone
	}

	// レアリティ別の重みとレアリティ別のアイテム一覧に分離
	weightByRarity := weightByRarities(gachaItems)
	itemsByRarity := itemsByRarities(gachaItems)

	random := CryptoRandomizer{}

	// 抽選処理
	results := make([]DrawResultItem, 0, count)
	for i := 0; i < count; i++ {
		// レアリティ抽選を実施
		rarity, err := drawRarity(random, weightByRarity)
		if err != nil {
			return nil, err
		}

		// 抽選対象アイテムからランダムに１件取得
		tempItems := itemsByRarity[rarity]
		if len(tempItems) == 0 {
			return nil, fmt.Errorf("gacha %d has no items of rarity %s", gachaID, rarity)
		}
		idx, err := random.Intn(len(tempItems))
		if err != nil {
			return nil, err
		}
		item := tempItems[idx]

		// 結果を返す
		results = append(results, DrawResultItem{
			ItemID:   item.ItemID,
			ItemName: item.ItemName,
			Rarity:   item.Rarity,
			IsPity:   false,
		})
	}

	if err := tans.Commit(); err != nil {
		return nil, err
	}

	return &DrawResult{Results: results, PityCount: 0}, nil
}