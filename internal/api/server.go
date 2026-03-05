package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() error {
	server := http.NewServeMux()

	server.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Printf("Listening on %s\n", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, server)
}
