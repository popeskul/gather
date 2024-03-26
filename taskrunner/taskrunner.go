package taskrunner

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run executes tasks from the tasks list in n parallel goroutines.
// concurrencyLevel - indicates the number of simultaneously running goroutines for executing tasks.
// errorLimit - indicates the number of errors after which the function will stop accepting new tasks and start the termination process.
func Run(tasks []Task, concurrencyLevel, errorLimit int) error {
	if concurrencyLevel <= 0 || len(tasks) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasksChan := make(chan Task, len(tasks))
	errChan := make(chan error, concurrencyLevel)
	var wg sync.WaitGroup

	// Goroutine for sending tasks to the tasksChan channel
	go func() {
		for _, task := range tasks {
			tasksChan <- task
		}
		close(tasksChan)
	}()

	// Goroutines for executing tasks
	for i := 0; i < concurrencyLevel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChan {
				if err := task(); err != nil {
					select {
					case errChan <- err:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	// Goroutine for handling errors
	go func() {
		var errCount int
		for err := range errChan {
			if err != nil {
				errCount++
				if errCount >= errorLimit {
					cancel()
					return
				}
			}
		}
	}()

	wg.Wait()
	close(errChan)

	if errorLimit > 0 && errors.Is(ctx.Err(), context.Canceled) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
