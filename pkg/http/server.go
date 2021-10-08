package http

import (
	"net/http"
)

type Server struct {
	Addr    string
	Scripts string
}

func NewServer(addr string, path string) *Server {
	return &Server{
		Addr:    addr,
		Scripts: path,
	}
}

func (server *Server) Serve() error {
	http.Handle("/", http.FileServer(http.Dir(server.Scripts)))

	return http.ListenAndServe(server.Addr, nil)
}
