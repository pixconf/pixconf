package xkit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRandomTime(t *testing.T) {
	for i := 0; i < 1000; i++ {
		randTime := RandomTime(10)

		require.True(t, randTime >= 0)
		require.True(t, randTime < 10*time.Second)
	}
}
