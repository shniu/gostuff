package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func handleCache(w http.ResponseWriter, req *http.Request) {

	path := req.URL.EscapedPath()
	key := strings.Split(path, "/")[3]
	fmt.Println("URL Path: ", path, key)

	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Method == http.MethodPut {
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println("body:", string(body))
		_, _ = w.Write([]byte("{\"code\":0,\"message\":\"succeed\"}"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func CreateCacheServer() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/api/cache/", handleCache)
	http.ListenAndServe(":5000", serveMux)
}
