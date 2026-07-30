package model

import "time"

type GachaHistory struct {
	ID         uint64
	UserID     uint64
	GachaID    uint64
	ItemID     uint64
	IsPity     bool
	DrawnAt    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
