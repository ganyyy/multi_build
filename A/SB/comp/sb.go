package comp

import (
	"math/rand/v2"
	"strconv"
)

func GenRandomSuffix() string {
	return strconv.Itoa(rand.IntN(1000))
}
