package csp

import (
	"testing"
	"math/rand"
	"time"
	"sync"
	"fmt"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var wg sync.WaitGroup

func f(s string) {
	defer wg.Done()
	// Sleep up to half a second
	delay := time.Duration(r.Int() % 500) * time.Millisecond
	time.Sleep(delay)
	fmt.Println(s)
}

func TestGoroutines01(t *testing.T) {
	fmt.Println("--- Run sequentially as normal functions")
	for i := 0; i < 4; i++ {
		wg.Add(3)
		f("1")
		f("2")
		f("3")
	}

	fmt.Println("--- Run concurrently as goroutines")
	for i := 0; i < 5; i++ {
		wg.Add(3)
		go f("1")
		go f("2")
		go f("3")
	}

	wg.Wait()
	fmt.Println("Done")
}
