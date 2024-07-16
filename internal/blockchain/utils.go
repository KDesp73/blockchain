package blockchain

import (
	"math/rand"
)
func random(min int, max int) int {
    return rand.Intn(max-min) + min
}
