package server

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Test struct {

}

func (t *Test) Echo() {
	fmt.Print("reply to client")
}

//
func NewServer() {
	addr := &net.TCPAddr{IP: []byte(""), Port: 9999}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Error listening", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on 9999")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}