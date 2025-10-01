package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed ui/dist/*
var embeddedUI embed.FS

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start() error {
	fsys, err := fs.Sub(embeddedUI, "ui/dist")
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))

	log.Printf("frontend: starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}
