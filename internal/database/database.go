package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Jeff-Rowell/hpotter/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB     *gorm.DB
	config types.DBConfig
}

func NewDatabase(config types.DBConfig) (*Database, error) {
	db := &Database{
		config: config,
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
	dialector := postgres.Open(d.buildPostgresConnectionString())

	var err error
	var errSlice []error
	for range 10 {
		d.DB, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			errSlice = append(errSlice, err)
			time.Sleep(2 * time.Second)
		}
	}

	if len(errSlice) == 10 {
		return fmt.Errorf("failed to open database connection 10 times: %w", err)
	}

	log.Printf("connected to %s database", d.config.DBType)
	return nil
}

func (d *Database) buildPostgresConnectionString() string {
	host := "localhost"
	port := "5432"
	dbname := "hpotter-database"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, d.config.User, d.config.Password, dbname, port)

	return dsn
}

func (d *Database) migrate() error {
	return Migrate(d.DB)
}

func (d *Database) Write(record any) error {
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
