package utils

import (
	"math/rand"
	"time"
)

// Randint 返回一个在 min 和 max 之间的随机整数，包含 min，但不包含 max
func Randint(min int, max int) int {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min // 交换 min 和 max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
