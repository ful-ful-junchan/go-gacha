package service

import (
	"errors"

	"go-app/internal/model"
	"go-app/internal/repository"
)

// ErrZeroWeight はレアリティ毎のweight合計が0のときに返される。
var ErrZeroWeight = errors.New("total weight is zero")

// rarityOrder はレアリティの走査順を固定するためのもの。
// mapのイテレーション順はGoでは保証されないため、累積分布の計算やテストの再現性のために順序を固定する。
var rarityOrder = []model.Rarity{model.RaritySSR, model.RaritySR, model.RarityR}

// weightByRarities はレアリティごとのweight合計を求める
func weightByRarities(gachaItems []repository.GachaItemEntry) map[model.Rarity]uint64 {
	list := make(map[model.Rarity]uint64)
	for _, item := range gachaItems {
		list[item.Rarity] += item.Weight
	}
	return list
}

// 排出アイテムをレアリティ別にグルーピング
func itemsByRarities(gachaItems []repository.GachaItemEntry) map[model.Rarity][]repository.GachaItemEntry {
	list := make(map[model.Rarity][]repository.GachaItemEntry)
	for _, item := range gachaItems {
		list[item.Rarity] = append(list[item.Rarity], item)
	}
	return list
}

func sumAllWeights(weightByRarity map[model.Rarity]uint64) uint64 {
	var total uint64
	for _, w := range weightByRarity {
		total += w
	}
	return total
}

// drawRarity は各レアリティのweight合計に基づいて1つのレアリティを重み抽選
func drawRarity(rng Randomizer, weightByRarity map[model.Rarity]uint64) (model.Rarity, error) {
	total := sumAllWeights(weightByRarity)
	if total == 0 {
		return "", ErrZeroWeight
	}
	v, err := rng.Intn(int(total))
	if err != nil {
		return "", err
	}
	n := uint64(v)

	var cumulative uint64
	for _, rarity := range rarityOrder {
		cumulative += weightByRarity[rarity]
		if n < cumulative {
			return rarity, nil
		}
	}
	return model.RarityR, nil
}