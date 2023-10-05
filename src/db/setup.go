package db

import (
	"errors"
	"strings"
	"time"

	"github.com/audriusdai/eventing-api/config"
	"github.com/audriusdai/eventing-api/core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initial connection to db. Assign global package variable 'DB'.
func SetupDb() error {
	db, err := gorm.Open(
		postgres.Open(getDsn()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		return errors.Join(err, errors.New("failed to connect to database"))
	}

	// models in an order
	db.AutoMigrate(
		&model.Event{},
	)

	sqlDB, err := db.DB()

	if err != nil {
		return errors.Join(err, errors.New("failed invoke DB() for further connection pool opening"))
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	return nil
}

func getDsn() string {
	return strings.Join(
		[]string{
			"host=" + config.DB_HOSTNAME,
			"port=" + config.DB_PORT,
			"user=" + config.DB_USERNAME,
			"password=" + config.DB_PASSWORD,
			"dbname=" + config.DB_NAME,
			"search_path=" + config.DB_SCHEMA,
		},
		" ",
	)
}
