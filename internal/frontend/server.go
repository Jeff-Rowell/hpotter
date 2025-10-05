package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/Jeff-Rowell/hpotter/internal/frontend/api"
	"gorm.io/gorm"
)

//go:embed ui/dist/*
var embeddedUI embed.FS

type Server struct {
	addr string
	db   *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Start() error {
	// Initialize API handler
	apiHandler := api.NewHandler(s.db)

	// API routes
	http.HandleFunc("/api/connections", apiHandler.GetConnections)
	http.HandleFunc("/api/connection", apiHandler.GetConnection)
	http.HandleFunc("/api/geo-data", apiHandler.GetGeoData)
	http.HandleFunc("/api/credentials", apiHandler.GetCredentials)
	http.HandleFunc("/api/data", apiHandler.GetData)
	http.HandleFunc("/api/stats", apiHandler.GetStats)

	// Serve static files
	fsys, err := fs.Sub(embeddedUI, "ui/dist")
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))

	log.Printf("frontend: starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}
