package repository

import (
	"context"
	"database/sql"
	"errors"

	"go-app/internal/model"
)

var ErrNotFound = errors.New("not found")

func FindGachaByID(ctx context.Context, db Executor, gachaID uint64) (*model.Gacha, error) {
	row := db.QueryRowContext(ctx, `
		SELECT id, name, pity_threshold, starts_at, ends_at, created_at
		FROM gachas
		WHERE id = ?
	`, gachaID)

	var g model.Gacha
	err := row.Scan(&g.ID, &g.Name, &g.PityThreshold, &g.StartsAt, &g.EndsAt, &g.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &g, nil
}
