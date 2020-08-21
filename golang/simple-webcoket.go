package ws	// https://dev.to/davidnadejdin/simple-server-on-gorilla-websocket-52h7

import (
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Accepting all requests
    },
}

type Server struct {
    clients       map[*websocket.Conn]bool
    handleMessage func(message []byte) // New message handler
}

func StartServer(handleMessage func(message []byte)) *Server {
    server := Server{
        make(map[*websocket.Conn]bool),
        handleMessage,
    }

    http.HandleFunc("/", server.echo)
    go http.ListenAndServe(":8080", nil)

    return &server
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
    connection, _ := upgrader.Upgrade(w, r, nil)

    server.clients[connection] = true // Save the connection using it as a key

    for {
        mt, message, err := connection.ReadMessage()

        if err != nil || mt == websocket.CloseMessage {
            break // Exit the loop if the client tries to close the connection or the connection is interrupted
        }

        go server.handleMessage(message)
    }

    delete(server.clients, connection) // Removing the connection

    connection.Close()
}

func (server *Server) WriteMessage(message []byte) {
    for conn := range server.clients {
        conn.WriteMessage(websocket.TextMessage, message)
    }
}

package main

import (
    "fmt"
    "simple-webcoket/ws"
)

func main() {
    server := ws.StartServer(messageHandler)

    for {
        server.WriteMessage([]byte("Hello"))
    }
}

func messageHandler(message []byte) {
    fmt.Println(string(message))
}
