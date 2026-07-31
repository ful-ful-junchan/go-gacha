package service

import (
	"crypto/rand"
	"math/big"
	mathrand "math/rand/v2"
)

// Randomizer は乱数生成を抽象化する。テスト時に固定値を返す実装に差し替えられるようにするため。
type Randomizer interface {
	// Intn は [0, n) の範囲の整数を返す。エントロピー源の読み取り失敗などをエラーとして返す。
	Intn(n int) (int, error)
}

// CryptoRandomizer は crypto/rand を使った本番用の実装。
// 実際の抽選結果を左右するため、予測不能性が求められる draw に使う。
type CryptoRandomizer struct{}

func (CryptoRandomizer) Intn(n int) (int, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(v.Int64()), nil
}

// MathRandomizer は math/rand/v2 を使った実装。
// DBに保存しない統計シミュレーション(verify)専用。大量試行でも高速に動作する。
type MathRandomizer struct{}

func (MathRandomizer) Intn(n int) (int, error) {
	return mathrand.IntN(n), nil
}
