package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
)

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func userAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User API Handler."))
}

type Handler1 interface {
	Serve(w string, r int)
}

type HandlerFunc func(w string, r int)

func (f HandlerFunc) Serve(w string, r int) {
	f(w, r)
}

func demo(w string, r int) {
	fmt.Println(w, r)
}

func HandleFunc1(p string, handler func(string, int)) {
	var hf HandlerFunc
	// 类型转换，把 handler 函数转换成 HandlerFunc 类型
	hf = HandlerFunc(handler)
	hf.Serve("123", 30)
}

// 自定义的 func type 用法
func customTypeFuncUsage() {
	HandleFunc1("", demo)
}

// 最简单的 http server 实现
func simpleHttpServer() {
	fmt.Println("Simple HTTP Server ...")
	http.Handle("/", &helloHandler{})
	http.HandleFunc("/user", userAPIHandler)
	http.ListenAndServe(":1234", nil)
}

// 最简单的文件服务器的实现
func simpleFileServer() {
	fmt.Println("Simple File Server ...")
	http.ListenAndServe(":1234", http.FileServer(http.Dir(".")))
}

// 增加了路由策略的 http server
func httpServerWithRouter() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// root path
		w.Write([]byte("Homepage"))
	})
	mux.HandleFunc("/order/", orderServiceHandler)
	mux.Handle("/payment/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Payment Service Handler"))
	}))

	mux.HandleFunc("/round-head", func(w http.ResponseWriter, r *http.Request) {
		// Get headers
		var headerKey string
		header := r.Header
		for k, v := range header {
			fmt.Println(k, v)
			if headerKey == "" {
				headerKey = k
			} else {
				headerKey = fmt.Sprintf("%s; %s", headerKey, k)
			}

			for _, entry := range v {
				w.Header().Add(k, entry)
			}
		}

		// 读取环境变量中的 VERSION 配置
		version := os.Getenv("VERSION")
		w.Header().Add("VERSION", version)

		w.WriteHeader(200)

		fmt.Printf("Client IP: %s, Response status: %d, Request Header: %s \n",
			r.RemoteAddr, 200, headerKey)

		w.Write([]byte(headerKey))
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	port := os.Getenv("MY_SERVICE_PORT")
	if port == "" {
		port = "1234"
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Println("Starting at", addr)
	http.ListenAndServe(addr, mux)
}

func orderServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order Service Handler."))
}

func main() {
	// customTypeFuncUsage()
	// simpleHttpServer()
	// simpleFileServer()

	flag.Set("v", "4")
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.V(2).Info("Starting http server ...")

	httpServerWithRouter()
}
