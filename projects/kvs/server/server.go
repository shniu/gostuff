package server

import (
	"encoding/json"
	"fmt"
	"github.com/shniu/gostuff/projects/kvs/log"
	"github.com/shniu/gostuff/projects/kvs/storage"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Http Server Graceful shutdown: https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2 Good!!!
// https://astaxie.gitbooks.io/build-web-application-with-golang/zh/03.3.html
// https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf
// Wrap for HandleFunc: https://golang.org/doc/articles/wiki/final.go
// A simple golang web server with basic logging, tracing, health check, graceful shutdown and zero dependencies, as below:
//   https://gist.github.com/creack/4c00ee404f2d7bd5983382cc93af5147
//   https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7
// Go Http Server Intro: https://cizixs.com/2016/08/17/golang-http-server-side/  Good!!!

var logger = log.Logger

// Kvs server
type Server struct {
	mux *http.ServeMux
	kvs storage.Kvs
}

type Response struct {
	Status string
	Data   string
}

func New(options ...func(*Server)) *Server {
	s := &Server{mux: http.NewServeMux()}

	for _, f := range options {
		f(s)
	}

	// GET /get?key=abc
	s.mux.HandleFunc("/get", s.get)
	s.mux.HandleFunc("/set", s.set)
	s.mux.HandleFunc("/delete", s.del)

	return s
}

func Kvs(kvs storage.Kvs) func(*Server) {

	return func(s *Server) {
		s.kvs = kvs
	}

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// https://developer.mozilla.org/zh-CN/docs/Glossary/Response_header
	// w.Header().Set("Server", "example Go server")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	s.mux.ServeHTTP(w, r)
}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	// Get Method
	if r.Method == http.MethodGet {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(w, "Error: ", err)
			return
		}

		key := values.Get("key")
		if len(key) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(w, "Error: Wrong input key or not found")
			return
		}

		bytes := s.kvs.Get(key)
		logger.Infof("/get?key=%s, value=%s\n", key, bytes)
		content := fmt.Sprintf("{\"status\":\"succeed\",\"value\":\"%s\"}", string(bytes))
		_, _ = fmt.Fprint(w, content)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		logger.Info(w, "Error: Only GET accepted")
	}
}

//type SetBody struct {
//	key   string `json:key`
//	value []byte `json:value`
//}

func (s *Server) set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	log.Println(w, "Error: ", err)
		//}

		// log.Printf("---- %v", string(bodyBytes))

		// https://cizixs.com/2016/12/19/golang-json-guide/
		// json parse
		var body map[string]interface{}
		//var body SetBody
		err = json.Unmarshal(bodyBytes, &body)
		// err := json.NewDecoder(r.Body).Decode(&body)
		logger.Info("body=%v, err=%v\n", body, err)

		key := body["key"].(string)
		value := body["value"].(string)
		//key := body.key
		//value := body.value
		logger.Info("key=", key, "value=", value)

		b, err := s.kvs.Put(key, []byte(value))
		content := fmt.Sprintf("{\"status\":\"succeed\"}")
		if !b || err != nil {
			content = fmt.Sprintf("{\"status\":\"failed\"}")
		}
		logger.Info("Response: ", content)
		_, _ = fmt.Fprint(w, content)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(w, "Error: Only POST accepted")
	}
}

// del == delete
func (s *Server) del(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {

		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(w, "Error: ", err)
			return
		}

		key := values.Get("key")
		if len(key) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(w, "Error: Wrong input key or not found")
			return
		}

		b, err := s.kvs.Delete(key)
		content := fmt.Sprintf("{\"status\":\"succeed\"}")
		if !b || err != nil {
			content = fmt.Sprintf("{\"status\":\"failed\"}")
		}
		logger.Info("Response: ", content)
		_, _ = fmt.Fprint(w, content)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(w, "Error: Only DELETE accepted")
	}
}
