package model

import "time"

type Gacha struct {
	ID            uint64
	Name          string
	PityThreshold uint64
	StartsAt      time.Time
	EndsAt        time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
