package database

import (
	"time"

	"gorm.io/gorm"
)

type Connections struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	SourceAddress      string    `gorm:"type:inet" json:"source_address"`
	SourcePort         int       `json:"source_port"`
	DestinationAddress string    `gorm:"type:inet" json:"destination_address"`
	DestinationPort    int       `json:"destination_port"`
	Latitude           float32   `json:"latitude"`
	Longitude          float32   `json:"longitude"`
	Container          string    `json:"container"`
	Proto              int       `json:"proto"`

	// Relationships
	Credentials []Credentials `gorm:"foreignKey:ConnectionsID" json:"credentials,omitempty"`
	Data        []Data        `gorm:"foreignKey:ConnectionsID" json:"data,omitempty"`
}

type Credentials struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	Username      string      `json:"username"`
	Password      string      `json:"password"`
	ConnectionsID uint        `json:"connections_id"`
	Connection    Connections `gorm:"foreignKey:ConnectionsID" json:"connection,omitempty"`
}
type Data struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Direction     string `json:"direction"`
	Data          string `json:"data"`
	ConnectionsID uint   `json:"connections_id"`

	// Relationship
	Connection Connections `gorm:"foreignKey:ConnectionsID" json:"connection,omitempty"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Connections{}, &Credentials{}, &Data{})
}
