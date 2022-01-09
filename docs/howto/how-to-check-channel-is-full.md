# How to check if the channel is full

via: https://stackoverflow.com/questions/25657207/how-to-know-a-buffered-channel-is-full

```go
package ch

import "fmt"

func isFull() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("Channel full. Discarding value")
    }
}

func isFull2() {
    ch := make(chan int, 1)
    ch <- 2
    if len(ch) == cap(ch) {
    	fmt.Println("Full")
    }
}
```
