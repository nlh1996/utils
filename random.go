package utils

import (
	"math/rand"
	"time"
)

// Rand .
func Rand(data map[uint8][]int) uint8 {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(10000)
	rand ++
	for k, v := range data {
		if rand >= v[0] && rand <= v[1] {
			return k
		}
	}
 	return 0
}
