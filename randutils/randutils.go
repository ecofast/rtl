package randutils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandRange(low, high int) int {
	if (low < 0) || (high < 0) || (low > high) {
		return 0
	}
	return low + rand.Intn(high - low + 1)
}
