package repository

import (
	"context"

	"go-app/internal/model"
)

// InsertGachaHistory は抽選履歴を1件保存する。
func InsertGachaHistory(ctx context.Context, db Executor, h model.GachaHistory) error {
	_, err := db.ExecContext(ctx, `
		INSERT INTO gacha_histories (user_id, gacha_id, item_id, is_pity, drawn_at, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, h.UserID, h.GachaID, h.ItemID, h.IsPity, h.DrawnAt, h.CreatedAt, h.UpdatedAt)
	return err
}
