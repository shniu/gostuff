package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const ServerDomain = "0.0.0.0:5001"
const ConnectDomain = "192.168.10.19:5001"

// 让两个人相互对话10万次
var totalCnt = uint32(100000)
var serverRecvCnt = uint32(0)
var clientRecvCnt = uint32(0)

// 并发写锁
var serverLock sync.Mutex
var clientLock sync.Mutex

// 传输协议
type RequestResponse struct {
	// 请求序号
	Serial uint32
	// 内容
	Payload []byte
}

// 异步设计

// 反序列化: 字节数组 -> 对象
func read(conn *net.TCPConn) *RequestResponse {
	// 读 4bytes 的总长度
	totalLengthBytes := make([]byte, 4)
	// conn.Read(totalLengthBytes)
	io.ReadFull(conn, totalLengthBytes)
	totalLength := binary.BigEndian.Uint32(totalLengthBytes)

	// 读 4bytes 的序号
	serialBytes := make([]byte, 4)
	// conn.Read(serialBytes)
	io.ReadFull(conn, serialBytes)
	serial := binary.BigEndian.Uint32(serialBytes)

	// 读 payload
	payloadBytes := make([]byte, totalLength-4)
	// conn.Read(payloadBytes)
	io.ReadFull(conn, payloadBytes)
	return &RequestResponse{Serial: serial, Payload: payloadBytes}
}

// 序列化: 对象 -> 字节数组
func write(conn *net.TCPConn, r *RequestResponse, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	// 4 bytes 表示总长度
	total := len(r.Payload) + 4
	totalLengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(totalLengthBytes, uint32(total))

	// 4 bytes 表示序号
	serialBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(serialBytes, r.Serial)

	// 剩下的全是Payload

	conn.Write(totalLengthBytes)
	conn.Write(serialBytes)
	conn.Write(r.Payload)
}

// 异步网络IO, 使用 TCP 通信
func startServer(wg *sync.WaitGroup) {
	fmt.Println("Start server")

	tcpAddr, err := net.ResolveTCPAddr("tcp", ServerDomain)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(tcpAddr)
	// Listen
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// 全双工通信
		// 接收客户端的消息
		go handleRecvMessageFromClient(tcpConn, wg)

		// 发送消息给客户端
		go handleSendMessageToClient(tcpConn)
	}
}

func handleSendMessageToClient(conn *net.TCPConn) {
	fmt.Println("向客户端的发送消息：", conn.LocalAddr().String(), conn.RemoteAddr().String())
	var nextSerial uint32
	for i := uint32(0); i < totalCnt; i++ {
		write(conn, &RequestResponse{Serial: nextSerial, Payload: []byte("I am server")}, &serverLock)
		nextSerial++
	}
}

func handleRecvMessageFromClient(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("接收来自客户端的消息：", conn.LocalAddr().String(), conn.RemoteAddr().String())
	for serverRecvCnt < totalCnt*2 {
		r := read(conn)
		// fmt.Println("Serial is ", r.Serial, " payload is ", string(r.Payload))
		go write(conn, &RequestResponse{Serial: r.Serial, Payload: []byte("TO Client")}, &serverLock)

		serverRecvCnt++
	}
	fmt.Println("Server recv: ", serverRecvCnt)
}

func startClient(wg *sync.WaitGroup) (*net.TCPConn, error) {
	fmt.Println("Start client")

	tcpAddr, err := net.ResolveTCPAddr("tcp", ConnectDomain)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// conn to server
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// 全双工通信
	// 接收服务端的消息
	go handleRecvMessageFromServer(tcpConn, wg)

	// 发送消息给服务端
	go handleSendMessageToServer(tcpConn)

	return tcpConn, nil
}

func handleSendMessageToServer(conn *net.TCPConn) {
	fmt.Println("向服务端的发送消息：", conn.LocalAddr().String(), conn.RemoteAddr().String())
	var nextSerial uint32
	for i := uint32(0); i < totalCnt; i++ {
		write(conn, &RequestResponse{Serial: nextSerial, Payload: []byte("hello")}, &clientLock)
		nextSerial++
	}
}

func handleRecvMessageFromServer(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("接收来自服务端的消息：", conn.LocalAddr().String(), conn.RemoteAddr().String())
	for clientRecvCnt < totalCnt*2 {
		r := read(conn)
		// fmt.Println("Serial is ", r.Serial, " payload is ", string(r.Payload))
		go write(conn, &RequestResponse{Serial: r.Serial, Payload: []byte("TO Server")}, &clientLock)
		clientRecvCnt++
	}
	fmt.Println("Client recv: ", clientRecvCnt)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动一个服务端
	go startServer(&wg)

	// 等待 Server 启动成功
	time.Sleep(time.Second)

	// 启动一个客户端
	conn, _ := startClient(&wg)
	// go startClient()

	// 暂时这么写，等待服务端和客户端逻辑处理完成
	// time.Sleep(time.Second * 10)
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	conn.Close()
	fmt.Println(" 耗时：", elapsed)
}
