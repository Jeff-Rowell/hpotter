package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jeff-Rowell/hpotter/internal/database"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// GetConnections returns all connections with optional pagination
func (h *Handler) GetConnections(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var connections []database.Connections
	query := h.db.Model(&database.Connections{})

	// Pagination
	if limit := r.URL.Query().Get("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(l)
		}
	}

	if offset := r.URL.Query().Get("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			query = query.Offset(o)
		}
	}

	// Order by most recent
	query = query.Order("created_at DESC")

	if err := query.Find(&connections).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connections)
}

// GetConnection returns a single connection with all related data
func (h *Handler) GetConnection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var connection database.Connections
	if err := h.db.Preload("Credentials").Preload("Data").First(&connection, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Connection not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connection)
}

// GetGeoData returns connections with geo-location data for map visualization
func (h *Handler) GetGeoData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var connections []database.Connections
	query := h.db.Model(&database.Connections{}).
		Where("latitude != 0 AND longitude != 0").
		Order("created_at DESC")

	// Optional limit for performance
	if limit := r.URL.Query().Get("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(l)
		}
	} else {
		query = query.Limit(1000) // Default limit
	}

	if err := query.Find(&connections).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connections)
}

// GetCredentials returns all credentials
func (h *Handler) GetCredentials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials []database.Credentials
	query := h.db.Model(&database.Credentials{})

	if limit := r.URL.Query().Get("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(l)
		}
	}

	if err := query.Find(&credentials).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(credentials)
}

// GetData returns payload data for a specific connection
func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	connectionID := r.URL.Query().Get("connection_id")
	if connectionID == "" {
		http.Error(w, "Missing connection_id parameter", http.StatusBadRequest)
		return
	}

	var data []database.Data
	if err := h.db.Where("connections_id = ?", connectionID).Find(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetStats returns aggregate statistics
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type Stats struct {
		TotalConnections int64            `json:"total_connections"`
		TotalCredentials int64            `json:"total_credentials"`
		TotalPayloads    int64            `json:"total_payloads"`
		UniqueIPs        int64            `json:"unique_ips"`
		TopCountries     []map[string]any `json:"top_countries"`
	}

	var stats Stats
	h.db.Model(&database.Connections{}).Count(&stats.TotalConnections)
	h.db.Model(&database.Credentials{}).Count(&stats.TotalCredentials)
	h.db.Model(&database.Data{}).Count(&stats.TotalPayloads)
	h.db.Model(&database.Connections{}).Distinct("source_address").Count(&stats.UniqueIPs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
