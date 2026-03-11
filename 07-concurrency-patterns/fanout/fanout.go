package fanout

import "sync"

func FanOut(input <-chan int, workers int, output chan<-int, wg *sync.WaitGroup) {
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range input {
				output <- value * value
			}
		}()
	}
}
