package xkit

import (
	"sync"
	"testing"
	"time"
)

func TestWaitTimeout(t *testing.T) {
	t.Run("CompletedNormally", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}()

		timeout := 200 * time.Millisecond
		if WaitTimeout(&wg, timeout) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("TimedOut", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			time.Sleep(200 * time.Millisecond)
			wg.Done()
		}()

		timeout := 100 * time.Millisecond
		if !WaitTimeout(&wg, timeout) {
			t.Errorf("Expected true, got false")
		}
	})
}
