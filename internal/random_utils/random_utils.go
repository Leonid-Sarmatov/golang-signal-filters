package random_utils

import (
	"math/big"
	"crypto/rand"
)

func RandomNumber(min, max int) int {
	rangeSize := big.NewInt(int64(max - min + 1))
	res, _ := rand.Int(rand.Reader, rangeSize)
	return min + int(res.Int64())
}