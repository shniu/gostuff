package csp

import (
	"fmt"
	"sync"
	"time"
)

func fib(c, quit chan int) {
	x, y := 0, 1

	for {
		c <- x
		select {
		case <-c:
			x, y = y, x+y
			// fmt.Println("---")
		case <-quit:
			fmt.Println("Quit")
			break
		}
	}

	// fmt.Println("x, y =", x, y)
}

func selectExample1() {
	c, quit := make(chan int, 1), make(chan int, 1)

	fib(c, quit)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(quit chan int) {
		time.Sleep(1 * time.Second)
		quit <- 0

		wg.Done()
	}(quit)

	wg.Wait()
}
