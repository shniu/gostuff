package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleCache(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println("body:", string(body))
		_, _ = w.Write([]byte("{\"code\":0,\"message\":\"succeed\"}"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func CreateCacheServer() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/api/cache", handleCache)
	http.ListenAndServe(":5000", serveMux)
}
