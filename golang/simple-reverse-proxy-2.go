package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var app appEnv
	err := app.parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "initialization error: %v\n", err)
		os.Exit(2)
	}
	if err = app.exec(); err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		os.Exit(1)
	}
}

type appEnv struct {
	URL      *url.URL
	ListenOn string
	FileName string
}

func (app *appEnv) parseArgs(args []string) error {
	fs := flag.NewFlagSet("simple-reverse-proxy", flag.ContinueOnError)
	proxyUrl := fs.String("proxy", "http://localhost:8000/", "URL to proxy requests to")
	listeningPort := fs.Int("listening-port", 80, "Port to listen on")
	allowExternal := fs.Bool("allow-external-connections", false, "Allow other computers to connect to your HTTP server")
	if err := fs.Parse(args); err != nil {
		return err
	}

	var err error
	app.URL, err = url.Parse(*proxyUrl)
	if err != nil {
		fs.Usage()
		return err
	}

	if app.URL.Scheme == "unix" {
		app.FileName = app.URL.Host
		// Handle URLs written like unix:whatever instead of unix://whatever
		if app.FileName == "" {
			app.FileName = app.URL.Opaque
		}
		app.URL = &url.URL{
			Scheme: "http",
			Host:   "127.0.0.1",
		}
	}

	if *allowExternal {
		app.ListenOn = fmt.Sprintf(":%d", *listeningPort)
	} else {
		app.ListenOn = fmt.Sprintf("127.0.0.1:%d", *listeningPort)
	}
	return nil
}

func LoggingMiddleware(s http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Reverse proxying %s ...\n", r.URL.Path)
		s.ServeHTTP(w, r)
	})
}

func (app *appEnv) SocketDialer(network, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", app.FileName)
}

func (app *appEnv) exec() error {
	proxied := app.URL.String()
	rp := httputil.NewSingleHostReverseProxy(app.URL)

	// Handle unix socket connections
	if app.FileName != "" {
		proxied = app.FileName
		rp.Transport = &http.Transport{
			Dial: app.SocketDialer,
		}
	}

	fmt.Printf("Started simple reverse proxy from %s to %s...\n\n",
		app.ListenOn, proxied)

	if err := http.ListenAndServe(app.ListenOn, LoggingMiddleware(rp)); err != nil {
		return err
	}
	return nil
}
