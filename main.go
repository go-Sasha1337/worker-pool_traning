package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
	"workerPool/miner"
)

func main() {
	var coal atomic.Int64
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("рабочий день воркеров окончен")
		cancel()
	}()
	initTime := time.Now()
	CoalChan := miner.Pool(ctx, 20)

	for z := range CoalChan {
		coal.Add(int64(z))
	}
	fmt.Println("итого:", coal.Load())
	fmt.Println("Затраченное время:", time.Since(initTime))

}
