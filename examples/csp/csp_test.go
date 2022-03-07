package csp

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestStopSubroutine(t *testing.T) {
	StopSubroutine()
}

func TestStopSubroutineWithContext(t *testing.T) {
	StopSubroutineWithContext()
}

func TestProducerToConsumer(t *testing.T) {
	producerToConsumer()
}

// Test: select
func TestSelectExample(t *testing.T) {
	selectExample1()
}

// Test: context
func TestContextDeadline(t *testing.T) {
	contextDeadline()
}

func ListDir(dir string, fn func (string, error) error) {
	err := fn("", nil)
	if err != nil {
		fmt.Println("ListDir...")
	}
}

func TestGoroutine1(t *testing.T) {
	ListDir("", func(s string, err error) error {
		go func() {
			fmt.Printf("%s", s)
		}()
		return err
	})
}

/// Test two server port
func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func TestTwoServerPort(t *testing.T) {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		err := serve(":8080", nil, stop)
		t.Logf("8080... %v", err)
		//
		//ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		//defer cancel()
		//select {
		//case <-ctx.Done():
		//	done <- errors.New("timeout error")
		//}
		done <- err
	}()
	go func() {
		done <- serve(":8081", nil, stop)
		t.Logf("8081...")
	}()

	t.Log("start...")

	var stopped bool
	for i := 0; i < cap(done); i++ {
		t.Log("start...")
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}

		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
