package mutex_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func TestMutexToAvoidRaceConditions(t *testing.T) {
	t.Run("should increment counter with Mutex lock to avoid race conditions", func(t *testing.T) {
		var wg sync.WaitGroup

		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go increment(&wg)
		}

		wg.Wait()
		t.Log("counter:", counter)
		assert.Equal(t, 10000, counter)
	})
}
