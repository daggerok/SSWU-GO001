package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name string
	Duration time.Duration
	Result int
}

func TaskRun(task *Task) {
	fmt.Printf("Task '%s' started\n", task.Name)
	task.Result = int(task.Duration.Seconds() * 1000)
	time.Sleep(task.Duration)
}

func FanOutTasks(wg *sync.WaitGroup, results chan<- Task, tasks ...Task) {
	for _, task := range tasks {
		wg.Add(1)

		go func(task Task) {
			defer wg.Done()
			TaskRun(&task)
			results <- task
		}(task)
	}
}

func FanInTasks(results <-chan Task, size int) {
	for i := 0; i < size; i++ {
		task := <- results
		fmt.Printf("Task '%s' completed with result: %d\n", task.Name, task.Result)
	}
}

func RunTaskWithContext(task *Task, ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Printf("Task '%s' not started: %v\n", task.Name, ctx.Err())
		return
	}

	// all good, we can continue with scheduler memory allocation for timers and so on...
	fmt.Printf("Task '%s' started\n", task.Name)

	select {
	case <-ctx.Done():
		fmt.Printf("Task '%s' canceled due to %v\n", task.Name, ctx.Err())
	case <-time.After(task.Duration):
		task.Result = int(task.Duration.Seconds() * 1000)
		fmt.Printf("Task '%s' completed with result: %d\n", task.Name, task.Result)
	}
}

func TimeoutManager(timeout time.Duration, tasks ...Task) []Task {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	size := len(tasks)
	fanOutResults := make(chan Task, size)

	for _, task := range tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			RunTaskWithContext(&task, ctx)
			fanOutResults <- task
		}(task)
	}

	go func() {
		wg.Wait()
		close(fanOutResults)
	}()

	var fanInResults []Task
	for task := range fanOutResults {
		fanInResults = append(fanInResults, task)
	}

	return fanInResults
}

func main() {
	// steps 1, 2, 3 and 4
	var wg sync.WaitGroup
	tasks := []Task{
		{"Hello World", time.Second, 0},
		{"Hello Moon", 2 * time.Second, 0},
		{"Hello All", 3 * time.Second, 0},
	}
	size := len(tasks)
	results := make(chan Task, size)

	FanOutTasks(&wg, results, tasks...)
	go func() {
		wg.Wait()
		close(results)
	}()
	FanInTasks(results, size)

	// steps 5 and 6
	otherTasks := []Task{
		{"Small", time.Second, 0},
		{"Medium", 2 * time.Second, 0},
		{"Slow", 3 * time.Second, 0},
	}
	TimeoutManager(time.Second * 3, otherTasks...)


	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//ch <- 3 // Что произойдет здесь? // selected deadlock answer, but quiz said its wrong... TODO: report this
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//// output:
	//// fatal error: all goroutines are asleep - deadlock!
	////
	//// goroutine 1 [chan send]:
	//// main.main()
	////        /Users/maksim.kostromin/Documents/code/private/SSWU-GO001/08-concurrent-task-management/main.go:119 +0x1dc
	//// exit status 2

	//
	var wg2 sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg2.Add(1)
		go worker(i, &wg2)
	}

	wg2.Wait()
	fmt.Println("All workers finished")

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker2(ctx)
	time.Sleep(3 * time.Second)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
