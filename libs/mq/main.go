package main

import (
	"fmt"
	"github.com/shniu/gostuff/libs/mq/server"
)

func main() {
	fmt.Print("Server startup.")
	server.NewServer()
}
