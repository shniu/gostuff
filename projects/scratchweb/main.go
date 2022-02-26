package main

import (
	"github.com/shniu/gostuff/projects/scratchweb/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: ":8765",
	}

	server.ListenAndServe()
}
