package fanout_test

import (
	"sync"
	"testing"

	"github.com/daggerok/SSWU-GO001/07-concurrency-patterns/fanout"
	"github.com/stretchr/testify/assert"
)

func TestFanOut(t *testing.T) {
	t.Run("should distribute tasks across multiple workers", func(t *testing.T) {
		// given task channel
		taskCh := make(chan int, 10)

		// and
		for i := 0; i < 10; i++ {
			taskCh <- i
		}
		close(taskCh)

		// and
		nrOfWorkers := 3
		results := make(chan int, 10)
		var wg sync.WaitGroup

		// when schedule tasks
		fanout.FanOut(taskCh, nrOfWorkers, results, &wg)

		// and wait for all tasks completion
		go func() {
			wg.Wait()
			close(results)
		}()

		// then collect all from results
		var finalResults []int
		for res := range results {
			finalResults = append(finalResults, res)
		}

		// and
		assert.Len(t, finalResults, 10, "Should have processed all 10 tasks")
		assert.Contains(t, finalResults, 1, "Should contain the first task result")
		assert.Contains(t, finalResults, 4, "Should contain the second task result")
		assert.Contains(t, finalResults, 81, "Should contain the last task result")
	})
}
