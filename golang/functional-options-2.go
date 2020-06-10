package app

type Server struct {
	listener net.Listener
	timeout  time.Duration
	cert     *tls.Cert
}

func NewServer(add string, options ...func(*Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := Server{listener: l}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

const WithTimeout = func(srv *Server) {
	srv.timeout = 60 * time.Second
}

const WithTLS = func(srv *Server) {
	config := loadTLSConfig()
	srv.listener = tls.NewListener(srv.listener, &config)
}

//=====================================================================

package main

func main() {
	// srv, _ := app.NewServer("localhost") // default

	/**
	timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
	}
	
	tls := func(srv *Server) {
		config := loadTLSConfig()
		srv.listener = tls.NewListener(srv.listener, &config)
	}
	
	srv, _ := app.NewServer("localhost", timeout, tls)
	*/

	srv, _ := app.NewServer("localhost", app.WithTimeout, app.WithTLS)
}
