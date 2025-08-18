package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/Jeff-Rowell/hpotter/types"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB         *gorm.DB
	config     types.DBConfig
	lockNeeded bool
	mu         sync.Mutex
}

func NewDatabase(config types.DBConfig) (*Database, error) {
	db := &Database{
		config:     config,
		lockNeeded: false,
	}

	// Determine if locking is needed (for SQLite)
	if config.DBType == "sqlite" {
		db.lockNeeded = true
	}

	if err := db.connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

func (d *Database) connect() error {
	var dialector gorm.Dialector

	switch d.config.DBType {
	case "sqlite":
		dialector = sqlite.Open(d.buildSQLiteConnectionString())
	case "postgres", "postgresql":
		dialector = postgres.Open(d.buildPostgresConnectionString())
	default:
		return fmt.Errorf("unsupported database type: %s", d.config.DBType)
	}

	var err error
	d.DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	log.Printf("Connected to %s database", d.config.DBType)
	return nil
}

func (d *Database) buildSQLiteConnectionString() string {
	if d.config.Name == "" {
		return "hpotter.db"
	}
	return d.config.Name
}

func (d *Database) buildPostgresConnectionString() string {
	user := d.config.User
	if user == "" {
		user = "postgres"
	}

	host := d.config.Host
	if host == "" {
		host = "localhost"
	}

	port := d.config.Port
	if port == "" {
		port = "5432"
	}

	dbname := d.config.Name
	if dbname == "" {
		dbname = "hpotter"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, d.config.Password, dbname, port)

	return dsn
}

func (d *Database) migrate() error {
	return Migrate(d.DB)
}

func (d *Database) Write(record interface{}) error {
	if d.lockNeeded {
		d.mu.Lock()
		defer d.mu.Unlock()
	}

	return d.DB.Create(record).Error
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	log.Println("Closing database connection")
	return sqlDB.Close()
}
