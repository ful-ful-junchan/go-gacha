package model

import "time"

type Rarity string

const (
	RaritySSR Rarity = "SSR"
	RaritySR  Rarity = "SR"
	RarityR   Rarity = "R"
)

type Item struct {
	ID        uint64
	Name      string
	Rarity    Rarity
	CreatedAt time.Time
	UpdatedAt time.Time
}
