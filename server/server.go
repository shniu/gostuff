package server

import (
	"encoding/json"
	"fmt"
	cache2 "github.com/shniu/cache/cache"
	"io/ioutil"
	"net/http"
	"strings"
)

type Server struct {
	cache cache2.Cache
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s *Server) ListenAndServe() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/api/cache/", s.handleCache)
	_ = http.ListenAndServe(":5000", serveMux)
}

func (s *Server) handleCache(w http.ResponseWriter, req *http.Request) {

	path := req.URL.EscapedPath()
	key := strings.Split(path, "/")[3]
	fmt.Println(req.Method, path, "key is:", key)

	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := &Response{
		Code: 0,
	}

	if req.Method == http.MethodPut {
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println("body:", string(body))

		e := s.cache.Set(key, body)
		if e != nil {
			response.Code = -1
			response.Message = "failed"
		} else {
			response.Message = "succeed"
		}

		bytes, _ := json.Marshal(response)
		_, _ = w.Write(bytes)
	} else if req.Method == http.MethodGet {
		val, e := s.cache.Get(key)
		if e != nil {
			if e.Error() == "the key does not exist" {
				response.Code = 404
				response.Data = ""
			} else {
				response.Code = -1
				response.Message = "failed"
			}
		} else {
			response.Code = 0
			response.Data = string(val)
		}

		bytes, _ := json.Marshal(response)
		_, _ = w.Write(bytes)
	} else if req.Method == http.MethodDelete {
		e := s.cache.Del(key)
		if e != nil {
			response.Code = -1
			response.Message = "failed"
		}

		bytes, _ := json.Marshal(response)
		_, _ = w.Write(bytes)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func New(c cache2.Cache) {
	server := &Server{cache: c}
	server.ListenAndServe()
}
