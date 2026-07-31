package repository

import (
	"context"
	"database/sql"
	"errors"
)

// 現在の天井カウントを返す。レコードが無ければ0を返す
func GetPityCount(ctx context.Context, db Executor, userID, gachaID uint64) (uint64, error) {
	row := db.QueryRowContext(ctx, `
		SELECT count FROM user_pity_counters WHERE user_id = ? AND gacha_id = ? FOR UPDATE
	`, userID, gachaID)

	var count uint64
	err := row.Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 天井カウントを保存する(無ければ作成、あれば更新)
func UpsertPityCount(ctx context.Context, db Executor, userID, gachaID, count uint64) error {
	_, err := db.ExecContext(ctx, `
		INSERT INTO user_pity_counters (user_id, gacha_id, count)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE count = VALUES(count)
	`, userID, gachaID, count)
	return err
}
