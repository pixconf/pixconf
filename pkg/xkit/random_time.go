package xkit

import (
	"math/rand/v2"
	"time"
)

// RandomTime returns a random time.Duration in range [0, max)
func RandomTime(maxTime float64) time.Duration {
	return time.Duration(rand.Float64()*maxTime) * time.Second
}
