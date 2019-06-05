package algorithms

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FisherYatesShuffle(bs []int) {
	for i := len(bs) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		tmp := bs[i]
		bs[i] = bs[j]
		bs[j] = tmp
	}
}
