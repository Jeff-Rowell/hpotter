package frontend

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

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

func (s *Server) Start(ctx context.Context) error {
	// Initialize API handler
	apiHandler := api.NewHandler(s.db)

	// Create a new ServeMux to avoid conflicts with default ServeMux
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/connections", apiHandler.GetConnections)
	mux.HandleFunc("/api/connection", apiHandler.GetConnection)
	mux.HandleFunc("/api/geo-data", apiHandler.GetGeoData)
	mux.HandleFunc("/api/credentials", apiHandler.GetCredentials)
	mux.HandleFunc("/api/data", apiHandler.GetData)
	mux.HandleFunc("/api/stats", apiHandler.GetStats)

	// Serve static files
	fsys, err := fs.Sub(embeddedUI, "ui/dist")
	if err != nil {
		return err
	}

	mux.Handle("/", http.FileServer(http.FS(fsys)))

	// Create server with graceful shutdown support
	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	// Channel to capture server errors
	serverErrors := make(chan error, 1)

	// Start server in goroutine
	go func() {
		log.Printf("frontend: starting server on %s", s.addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Wait for context cancellation or server error
	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			return err
		}
	case <-ctx.Done():
		log.Printf("frontend: received shutdown signal")

		// Create shutdown context with timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt graceful shutdown
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("frontend: error during shutdown: %v", err)
			return err
		}
		log.Printf("frontend: server stopped gracefully")
	}

	return nil
}
