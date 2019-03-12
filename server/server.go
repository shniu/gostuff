package server

import (
	"fmt"
	cache2 "github.com/shniu/cache/cache"
	"io/ioutil"
	"net/http"
	"strings"
)

type Server struct {
	cache cache2.Cache
}

func (s *Server) ListenAndServe() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/api/cache/", s.handleCache)
	http.ListenAndServe(":5000", serveMux)
}

func (s *Server) handleCache(w http.ResponseWriter, req *http.Request) {

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
		e := s.cache.Set(key, body)
		if e != nil {
			_, _ = w.Write([]byte("{\"code\":-1,\"message\":\"failed\"}"))
		} else {
			_, _ = w.Write([]byte("{\"code\":0,\"message\":\"succeed\"}"))
		}
	} else if req.Method == http.MethodGet {
		fmt.Println("get value of ", key)
		val, e := s.cache.Get(key)
		if e != nil {
			_, _ = w.Write([]byte("{\"code\":-1,\"message\":\"failed\"}"))
		} else {
			_, _ = w.Write([]byte(fmt.Sprintf("{\"code\":0,\"data\":\"%v\"}", string(val))))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func New(c cache2.Cache) {
	server := &Server{cache: c}
	server.ListenAndServe()
}
