package csp

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Stop sub routine
func StopSubroutine() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt ...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(time.Second)
	fmt.Println("main process exit!")
}

func StopSubroutineWithContext() {
	baseCtx := context.Background()

	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
}

// Simple producer-consumer pattern
var queue = make(chan int, 10)

func producerToConsumer() {
	// queue := make(chan int, 10)

	// producer
	go func(queue chan<- int) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			if len(queue) != cap(queue) {
				queue<-100
			}
		}
	}(queue)

	// consumer
	go func(queue <-chan int) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			elem := <-queue
			fmt.Printf("Received %d\n", elem)
		}
	}(queue)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
