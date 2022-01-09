package main

import (
	"sync"
	"time"
	"net"
	"fmt"
	"encoding/binary"
	"io"
	"strconv"
)

// 服务端接收到数据的次数
var serverRecvCnt = uint32(0)
// 客户端收到数据的次数
var clientRecvCnt = uint32(0)
// 总共需要发送数据的次数
var total = uint32(100000)

var s0 = " 吃了没，您呐？"
var s3 = " 嗨！吃饱了溜溜弯儿。"
var s5 = " 回头去给老太太请安！"
var c1 = " 刚吃。"
var c2 = " 您这，嘛去？"
var c4 = " 有空家里坐坐啊。"

// 服务端写锁
var serverWriteLock sync.Mutex
// 客户端写锁
var clientWriteLock sync.Mutex

type RequestResponse struct {
	Serial  uint32 // 序号
	Payload string // 内容
}

// 接收数据，反序列化成 RequestResponse
func readFrom(conn *net.TCPConn) (*RequestResponse, error) {
	ret := &RequestResponse{}
	buf := make([]byte, 4)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf(" 读取长度错误：%s", err.Error())
	}
	length := binary.BigEndian.Uint32(buf)

	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf(" 读取长度错误：%s", err.Error())
	}

	ret.Serial = binary.BigEndian.Uint32(buf)
	payloadBytes := make([]byte, length-4)
	if _, err := io.ReadFull(conn, payloadBytes); err != nil {
		return nil, fmt.Errorf(" 读取长度错误：%s", err.Error())
	}
	ret.Payload = string(payloadBytes)
	// fmt.Println("接收到：" + strconv.Itoa(int(ret.Serial)) + ret.Payload)
	return ret, nil
}

// 序列化 RequestResponse
// 长度    serial长度    payload
// 4Byte   4Byte        变长
func writeTo(r *RequestResponse, conn *net.TCPConn, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	payloadBytes := []byte(r.Payload)
	// serial 长度
	serialBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(serialBytes, r.Serial)
	// 计算总长度
	length := uint32(len(payloadBytes) + len(serialBytes))
	lengthByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthByte, length)

	// 写入总长度
	conn.Write(lengthByte)
	// 写入serial
	conn.Write(serialBytes)
	// 写入payload
	conn.Write(payloadBytes)
	// fmt.Println(" 发送" + r.Payload)
}

func startServer(wg *sync.WaitGroup) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println(" Server 端等待客户端连接...")
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(" 客户端连接一次: " + conn.RemoteAddr().String())
		go serverListen(conn, wg)
		go serverSay(conn)
	}
}

func serverSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, s0}, conn, &serverWriteLock)
		nextSerial++
	}
	fmt.Println("服务端发送数据次数：" + strconv.Itoa(int(nextSerial)))
}

func serverListen(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()

	for serverRecvCnt < total*3 {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if r.Payload == c2 {
			// 收到：你这，嘛去？回复：嗨！吃饱了溜溜弯儿。
			go writeTo(&RequestResponse{r.Serial, s3}, conn, &serverWriteLock)
		} else if r.Payload == c4 {
			// 收到：有空家里坐坐啊。回复：回头给老太太请安！
			go writeTo(&RequestResponse{r.Serial, s5}, conn, &serverWriteLock)
		} else if r.Payload == c1 {
			// 收到：刚吃。不用回复
		} else {
			fmt.Println(" 服务端听不懂：" + r.Payload)
			break
		}

		serverRecvCnt++
	}
	fmt.Println("服务端接收到数据次数：" + strconv.Itoa(int(serverRecvCnt)))
}

func startClient(wg *sync.WaitGroup) *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	go clientListen(conn, wg)
	go clientSay(conn)
	return conn
}

func clientSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, c2}, conn, &clientWriteLock)
		nextSerial++
		writeTo(&RequestResponse{nextSerial, c4}, conn, &clientWriteLock)
		nextSerial++
	}
	fmt.Println("客户端发送数据次数：" + strconv.Itoa(int(nextSerial)))
}

func clientListen(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()

	for clientRecvCnt < total*3 {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if r.Payload == s0 {
			writeTo(&RequestResponse{r.Serial, c1}, conn, &clientWriteLock)
		} else if r.Payload == s3 {
			// do nothing
		} else if r.Payload == s5 {
			// do nothing
		} else {
			fmt.Println(" 客户端听不懂：" + r.Payload)
			break
		}

		clientRecvCnt++
		//fmt.Println("客户端接收到数据次数：" + strconv.Itoa(int(clientRecvCnt)))
	}
	fmt.Println("客户端接收到数据次数：" + strconv.Itoa(int(clientRecvCnt)))
}

func main() {
	// 启动服务端
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(&wg)
	time.Sleep(time.Second)

	// 启动客户端
	conn := startClient(&wg)
	t1 := time.Now()
	wg.Wait()
	elapsed := time.Since(t1)
	conn.Close()
	fmt.Println("耗时：", elapsed)
}
