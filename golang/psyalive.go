package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
)

var logErr = log.New(os.Stderr, "", 0)

func main() {

	parser := argparse.NewParser("", "Test if a server is alive by a given port")
	s := parser.String("s", "server", &argparse.Options{Required: true, Help: "Server to check"})
	p := parser.String("p", "port", &argparse.Options{Required: true, Help: "TCP Port to check"})
	n := parser.String("n", "notify", &argparse.Options{Required: false, 
		Help: "Service to send notifications. Currently is only stderr and telegram available. 
			   If not set it will notify over stderr. For telegram please use the syntax: 
			   tg://BOTTOKEN/ID"})
	t := parser.Int("t", "timeout", &argparse.Options{Required: false, 
		Help: "Timeout in Seconds. Default is 3min = 180"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	timeout := *t
	port := *p
	server := *s
	notification := *n

	if timeout == 0 {
		timeout = 180
	}

	if server == "" || port == "" {
		os.Exit(1)
	}
	alive := aliveTest(server, port, timeout)

	if !alive {
		if notification == "" {
			sendNotification("stderr", server, port)
		} else {
			sendNotification(notification, server, port)
		}
	}
}

func aliveTest(server string, port string, timeout int) bool {
	ip, err := net.LookupIP(server)
	if err != nil {
		logErr.Println(err)
	} else {
		address := fmt.Sprintf("%s:%s", ip[0], port)
		timeout := net.Dialer{Timeout: time.Duration(timeout) * time.Second}
		conn, err := timeout.Dial("tcp", address)
		if err != nil {
			return false
		} else {
			conn.Close()
			return true
		}

	}
	return false
}

func sendNotification(notification string, server string, port string) {
	if notification == "stderr" {
		logErr.Printf("Warning:\nThe Server %s is unavailable!\nTested Port: %s\n", server, port)
	} else {
		splits := strings.Split(notification, "://")
		if splits[0] == "tg" {
			telegram := strings.Split(splits[1], "/")
			URI := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?
				chat_id=%s&parse_mode=HTML&text=<b>%%E2%%9A%%A0Warning%%0a</b>
				The Server %s is unavailable!%%0aTested Port: %s", 
				telegram[0], telegram[1], server, port)
			http.Get(URI)
		} else {
			fmt.Println("Unsupported Notification Parameter! Using stderr.")
			sendNotification("stderr", server, port)
		}
	}
}

/* Notes:
emojis in URI: https://github.com/gasparandr/emoji-express/blob/master/data/emojis.js
*/
