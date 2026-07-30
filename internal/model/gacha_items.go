package model

import "time"

type GachaItem struct {
	GachaID   uint64
	ItemID    uint64
	Weight    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
