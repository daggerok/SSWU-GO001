package fanin_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/07-concurrency-patterns/fanin"
	"github.com/stretchr/testify/assert"
)

func FanInTest(t *testing.T) {
	t.Run("should merge results from multiple channels into one", func(t *testing.T) {
		// given 2 channels
		ch1 := make(chan string)
		ch2 := make(chan string)

		// when merge
		merged := fanin.FanIn(ch1, ch2)

		// and send data
		go func() {
			ch1 <- "ch1: hello"
			ch2 <- "ch2: world"
			close(ch1)
			close(ch2)
		}()

		// then
		var results []string
		for i := 0; i < 2; i++ {
			results = append(results, <-merged)
		}

		// and
		assert.Contains(t, results, "ch1: hello")
		assert.Contains(t, results, "ch2: world")
		assert.Len(t, results, 2)
	})
}
