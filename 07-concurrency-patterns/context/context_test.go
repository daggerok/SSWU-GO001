package context_test

import (
	"context"
	"testing"
	"time"
)

func worker(t *testing.T, ctx context.Context) {
	for { // can be: for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			t.Log("Work cancelled")
			return
		default:
			t.Log("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestContextAndCancellation(t *testing.T) {
	t.Run("should cancel tasks across multiple workers", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
		defer cancel()

		go worker(t, ctx)
		time.Sleep(3 * time.Second) // allow timeout to expire
	})
}
