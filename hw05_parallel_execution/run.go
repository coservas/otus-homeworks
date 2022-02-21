package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	c := make(chan Task)
	var errCount int32
	wg := sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for task := range c {
				err := task()
				if err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCount) >= int32(m) {
			break
		}
		c <- task
	}
	close(c)
	wg.Wait()

	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
