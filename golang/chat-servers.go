/*
go语言渐入佳境-网络[17]-go语言建立聊天服务器的3种案例赏析
2019-01-19

案例1

如何在55行Golang中编写TCP聊天服务器? go net包允许你编写TCP服务器。这是一个聊天服务器，
客户端发送的每个字节都被复制到每个其他客户端（包括发送者）

代码比较精彩，在主程序中建立了3个通道，分别是新链接、断开链接、广播信息。作者非常巧妙
的将代码压缩到55行，

虽然每一个客户端都在抢夺通道的信息，但是作者通过通道的缓冲区来缓解这个问题，目前这套
代码已经能够处理非常大量的并发聊天。

此代码的缺陷在于，在不断的广播过程中，可能会开辟无数的协程处理写入数据的操作，这时会
造成消息堵塞，甚至消息不按照顺序到达。
*/
package main

import "net"

func main() {
	newConns := make(chan net.Conn, 128)//新链接
	deadConns := make(chan net.Conn, 128)//断开链接
	publishes := make(chan []byte, 128)//广播信息
	conns := make(map[net.Conn]bool)
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	if err != nil { panic(err) }
  //防止卡住
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil { panic(err) }
			newConns <- conn
		}
	}()
	for {
		select {
		case conn := <-newConns:
      //新建链接后，会开辟协程不断短期客户端发出的消息
			conns[conn] = true
			go func() {
				buf := make([]byte, 1024)
				for {
					nbyte, err := conn.Read(buf)
					if err != nil {
						deadConns <- conn
						break
					} else {
						fragment := make([]byte, nbyte)
						copy(fragment, buf[:nbyte])
						publishes <- fragment
					}
				}
			}()
		case deadConn := <-deadConns:
      //断开链接、关闭资源
			_ = deadConn.Close()
			delete(conns, deadConn)
      //缺陷
		case publish := <-publishes:
      // 广播给所有的
			for conn, _ := range conns {
				go func(conn net.Conn) {
					totalWritten := 0
					for totalWritten < len(publish) {
						writtenThisCall, err := conn.Write(publish[totalWritten:])
						if err != nil {
							deadConns <- conn
							break
						}
						totalWritten += writtenThisCall
					}
				}(conn)
			}
		}
	}
}
/*
案例二：

案例二是一个服务器与客户端都混合在一起的例子。

案例二相对于案例一最大的改进在于，服务器为每一个客户端都新建了唯一的协程来处理数据。
在协程中，通过一个通道来通信。当客户端通道接收到数据，即会往客户端发送消息。

程序更加稳健。
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//每一个客户端，服务器都新建了通道
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	socket net.Conn
	data   chan []byte
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (manager *ClientManager) receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
			manager.broadcast <- message
		}
	}
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}

func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

func startServerMode() {
	fmt.Println("Starting server...")
	listener, error := net.Listen("tcp", ":12345")
	if error != nil {
		fmt.Println(error)
	}
	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for {
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}

func startClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", "localhost:12345")
	if error != nil {
		fmt.Println(error)
	}
	client := &Client{socket: connection}
	go client.receive()
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		connection.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startServerMode()
	} else {
		startClientMode()
	}
}

/*
案例3:

案例三是go语言圣经中的一段代码demo。 这段代码有点意思，服务器端绑定的是一个通道
而不是socker连接的指针。

同时每一个客户端都有一个携程来处理发送信息的操作。从这一点和案例二很相似。但是
缺陷也很大，不能处理断开连接的操作。只能关闭通道。

和案例二是没有办法比较的。
*/

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
/*
总结

案例二具有良好的稳健型。服务器为每一个客户端都开辟了一个协程处理发送数据操作，通过通道
来进行通信。能够解决案例三中不能解决的断开网络连接的操作，也能够解决案例一中开辟无数个
协程以及消息堵塞不同步的问题。。
*/
