package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() error {
	server := http.NewServeMux()

	fmt.Println("Listening on " + s.listenAddr)

	server.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir(filepath.Join("web", "static"))),
		),
	)

	server.HandleFunc("/", handleHomeRender)

	return http.ListenAndServe(s.listenAddr, server)
}

func handleHomeRender(w http.ResponseWriter, r *http.Request) {
	data := new(struct {
		Content string
	})

	data.Content = "Content from server"

	tplPath := filepath.Join("web", "templates", "index.html")
	te, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := te.Execute(w, data); err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}
