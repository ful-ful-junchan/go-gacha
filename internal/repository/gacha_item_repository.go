package repository

import (
	"context"

	"go-app/internal/model"
)

type GachaItemEntry struct {
	ItemID   uint64
	ItemName string
	Rarity   model.Rarity
	Weight   uint64
}

func FindGachaItemsByGachaID(ctx context.Context, db Executor, gachaID uint64) ([]GachaItemEntry, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT gi.item_id, i.name, i.rarity, gi.weight
		FROM gacha_items gi
		JOIN items i ON i.id = gi.item_id
		WHERE gi.gacha_id = ?
	`, gachaID)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var gachaItems []GachaItemEntry
	for rows.Next() {
		var e GachaItemEntry
		if err := rows.Scan(&e.ItemID, &e.ItemName, &e.Rarity, &e.Weight); err != nil {
			return nil, err
		}
		gachaItems = append(gachaItems, e)
	}
	return gachaItems, rows.Err()
}
