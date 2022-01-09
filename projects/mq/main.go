package main

import (
	"fmt"
	"github.com/shniu/gostuff/projects/mq/server"
)

func main() {
	fmt.Print("Server startup.")
	server.NewServer()
}
