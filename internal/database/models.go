package database

import (
	"net"
	"time"

	"gorm.io/gorm"
)

// Protocol constants from the original Python code
const (
	TCP = 6
	UDP = 17
)

// Connections represents the schema for all connections made to HPotter
type Connections struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	SourceAddress      net.IP    `gorm:"type:inet" json:"source_address"`
	SourcePort         int       `json:"source_port"`
	DestinationAddress net.IP    `gorm:"type:inet" json:"destination_address"`
	DestinationPort    int       `json:"destination_port"`
	Latitude           string    `json:"latitude"`
	Longitude          string    `json:"longitude"`
	Container          string    `json:"container"`
	Proto              int       `json:"proto"`

	// Relationships
	Credentials []Credentials `gorm:"foreignKey:ConnectionsID" json:"credentials,omitempty"`
	Data        []Data        `gorm:"foreignKey:ConnectionsID" json:"data,omitempty"`
}

// Credentials stores username and passwords where appropriate
type Credentials struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	ConnectionsID uint   `json:"connections_id"`

	// Relationship
	Connection Connections `gorm:"foreignKey:ConnectionsID" json:"connection,omitempty"`
}

// Data represents the requests (and possibly responses) to/from HPotter containers
type Data struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Direction     string `json:"direction"`
	Data          string `json:"data"`
	ConnectionsID uint   `json:"connections_id"`

	// Relationship
	Connection Connections `gorm:"foreignKey:ConnectionsID" json:"connection,omitempty"`
}

// Migrate runs the database migrations for all models
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Connections{}, &Credentials{}, &Data{})
}
