package main

import (
	"fmt"
	"github.com/shniu/cache/cache"
	"github.com/shniu/cache/server"
)

func main() {
	fmt.Println("Start Cache server, and listen on 0.0.0.0:5000")
	c := cache.NewCache("inMemoryCache")
	server.New(c)
}
