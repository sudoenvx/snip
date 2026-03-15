package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sudoenvx/snip/internal/api/handlers"
	"github.com/sudoenvx/snip/internal/database"
)

type Server struct {
	listenAddr string
	Db *database.DB
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func(s *Server) SetupDatabase() {
	connection := os.Getenv("DATABASE_URL")
	db, err := database.NewDB(context.Background(), connection)

	if err != nil {
		log.Fatal(err)
	}

	s.Db = db
	println(s.Db.Pool.Config().ConnConfig.Database)
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

	server.HandleFunc("/", handlers.HandleHomeRender)
	server.HandleFunc("POST /shorten", handlers.CreateShortenUrlHandler(s.Db))
	server.HandleFunc("GET /e/{code}", handlers.CreateRedirectHandler(s.Db))
	server.HandleFunc("GET /shorten-urls", handlers.CreateGetAllUrlsHandler(s.Db))

	return http.ListenAndServe(s.listenAddr, server)
}
