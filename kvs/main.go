package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/shniu/gostuff/kvs/log"
	"github.com/shniu/gostuff/kvs/options"
	"github.com/shniu/gostuff/kvs/server"
	"github.com/shniu/gostuff/kvs/storage"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var logger = log.Logger

func main() {
	logger.Info("============ Start K V S ============")

	home := flag.String("home", "/tmp/kvs", "Home directory for kvs db")
	port := flag.String("port", "3000", "Server listen address")
	flag.Parse()
	logger.Info("Server config:")

	opts := options.Options
	if *home != "" && strings.Index(*home, "") >= 0 {
		opts.Home = *home
	}
	opts.Print()

	// Init and load kvs
	kvs := openKvs()

	// Start server
	kvServer := serverSetup(kvs, port)

	go func() {
		logger.Infof("Listening on http://0.0.0.0%s\n", kvServer.Addr)
		if e := kvServer.ListenAndServe(); e != http.ErrServerClosed {
			logger.Fatalln(e)
		}
	}()

	// Graceful shutdown
	graceful(kvServer, kvs, 5*time.Second)
}

// Open Kvsdb
func openKvs() storage.Kvs {
	kvs, err := storage.Open(options.Options)
	if err != nil {
		logger.Fatalf("Initialization of kvs failed!")
	}
	return kvs
}

func serverSetup(kvs storage.Kvs, port *string) *http.Server {

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = fmt.Sprintf(":%s", *port)
	}

	return &http.Server{
		Addr:    addr,
		Handler: server.New(server.Kvs(kvs)),
	}
}

func graceful(hs *http.Server, kvs storage.Kvs, timeout time.Duration) {

	// close
	defer kvs.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	logger.Infof("Shutdown with timeout: %s\n", timeout)
	if err := hs.Shutdown(ctx); err != nil {
		logger.Infof("Error: %v\n", err)
	} else {
		logger.Info("============ Server graceful stopped, bye bye ============")
	}
}
