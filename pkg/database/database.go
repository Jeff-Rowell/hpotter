package database

import (
	"log/slog"
	"time"

	"github.com/hpotter/pkg/dockerclient"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DbImage      string
	DbPort       int
	Username     string
	Password     string
	DockerClient *dockerclient.Docker
	containerId  string
	db           *gorm.DB
}

func (d *Database) Create() error {
	err := d.DockerClient.CreateDbVolume()
	if err != nil {
		return err
	}

	networkId, err := d.DockerClient.GetDbNetworkId()
	if err != nil {
		return err
	}

	if networkId == "" {
		networkId, err = d.DockerClient.CreateDbNetwork()
		if err != nil {
			return err
		}
	}

	cId, err := d.DockerClient.CreateDbContainer(d.Username, d.Password, d.DbImage, d.DbPort)
	if err != nil {
		return err
	}

	d.containerId = cId

	err = d.DockerClient.StartDbContainer(cId)
	if err != nil {
		e := d.DockerClient.ContainerStop(cId, true)
		if e != nil {
			slog.Error("stop container error", "error", e)
		}
		return err
	}

	cIp, err := d.DockerClient.GetDBContainerIP(cId)
	if err != nil {
		e := d.DockerClient.ContainerStop(cId, true)
		if e != nil {
			slog.Error("stop container error", "error", e)
		}
		return err
	}

	dsn := d.buildConnectionString(cIp)
	dialector := postgres.Open(dsn)

	var e error
	var db *gorm.DB

	for range 10 {
		db, e = gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if e == nil {
			break
		}

		e = nil
		time.Sleep(1 * time.Second)
	}

	if e != nil {
		err := d.DockerClient.ContainerStop(cId, true)
		if err != nil {
			slog.Error("stop container error", "error", e)
		}
		return e
	}

	err = db.AutoMigrate(&Connection{}, &Credential{}, &Payload{})
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *Database) HandleShutdown() {
	err := d.DockerClient.ContainerStop(d.containerId, true)
	if err != nil {
		slog.Error("failed to remove db container", "error", err)
	}
}

func (d *Database) buildConnectionString(cIp string) string {
	port := "5432"
	dsn := "host=" + cIp
	dsn += " user=" + d.Username
	dsn += " password=" + d.Password
	dsn += " dbname=" + d.DockerClient.GetDbName()
	dsn += " port=" + port
	dsn += " sslmode=disable"
	return dsn
}

func (d *Database) CreateConnection(c *Connection) (uint, error) {
	result := d.db.Create(c)
	if result.Error != nil {
		return 0, result.Error
	}
	return c.ID, nil
}

func (d *Database) CreateCredential(c *Credential) error {
	result := d.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database) CreatePayload(p *Payload) error {
	result := d.db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
