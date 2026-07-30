package model

import "time"

type UserPityCounter struct {
	UserID    uint64
	GachaID   uint64
	Count     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
