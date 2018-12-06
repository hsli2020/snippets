golang代码片段（摘抄）

以下是从golang并发编程实战2中摘抄过来的代码片段，主要是实现一个简单的tcp socket通讯
（客户端发送一个数字，服务端计算该数字的立方根然后返回），写的不错，用到了go的并发以
及看下郝林大神是如何处理socket通讯的。

package main

import (
	"net"
	"strings"
	"fmt"
	"time"
	"io"
	"bytes"
	"strconv"
	"math"
	"math/rand"
	"sync"
)
const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER = '\t'
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

func serverGo() {
	defer wg.Done()
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
		return
	}
	defer listener.Close()
	printServerLog("Got listener for the server.(local address: %s)", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept Error: %s", err)
		}
		printServerLog("Established a connection with a client application,(remote address: %s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func printLog(role string, sn int, format string, args ...interface{})  {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{})  {
	printLog("Server", 0, format, args...)
}

func printClientLog(sn int, format string, args ...interface{})  {
	printLog("Cient", sn, format, args...)
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(10 * time.Second))		// set read timeline 10s
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side.")
			} else {
				printServerLog("Read Error: %s", err)
			}
			break
		}
		printServerLog("Received request: %s", strReq)
		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("Sent error message (written %d bytes): %s", n, err)
			continue
		}
		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("The cube root of %d is %f.", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("Sent response message (written %d bytes: %s).", n, respMsg)
		}
	}
}

func read(conn net.Conn) (string, error){
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func strToInt32(str string) (int32, error){
	i32, err := strconv.ParseInt(str, 10 ,32)
	return int32(i32), err
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func cbrt(intReq int32) float64 {
	return math.Cbrt(float64(intReq))
}

func clientGo(id int)  {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2 * time.Second)
	if err != nil {
		printClientLog(id, "Dial Error: %s", err)
		return
	}
	defer conn.Close()
	printClientLog(id, "Connected to server.(remote address: %s,local address: %s)", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i< requestNumber; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "Write Error: %s", err)
			continue
		}
		printClientLog(id, "Sent request (written %d bytes)", n)
	}
	for j := 0; j< requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientLog(id, "The connection is closed by another side.")
			}else{
				printClientLog(id, "Read Error: %s", err)
			}
			break
		}
		printClientLog(id, "Received response: %s", strResp)
	}
}
