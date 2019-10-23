package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("Client startup.")
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connecting to localhost:9999")

	// 使用 WaitGroup 进行优化
	// done := make(chan string)
	// go handleWrite(conn, done)
	// go handleRead(conn, done)

	// fmt.Println(<-done)
	// fmt.Println(<-done)

	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

// func handleRead(conn net.Conn, done chan string) {
func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	// done <- "Read"
}

// func handleWrite(conn net.Conn, done chan string) {
func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		_, e := conn.Write([]byte("hello" + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}
	// done <- "Sent"
}
