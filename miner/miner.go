package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int, result chan<- int, id int) {
	defer wg.Done()
	fmt.Println("Я шахтёр номер:", id, "Начал добывать уголь!")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", id, "Мой рабочий день закончен!")
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			result <- j * 3
			time.Sleep(time.Second)
		}
	}
}

func Pool(ctx context.Context, miner int) <-chan int {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	wg := sync.WaitGroup{}

	for w := 1; w <= miner; w++ {
		wg.Add(1)
		go Worker(ctx, &wg, jobs, result, w)
	}

	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(result)
	}()
	return result

}
