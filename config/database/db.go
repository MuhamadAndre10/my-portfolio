package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	var (
		env      = os.Getenv("ENV")
		host     = os.Getenv("PGHOST")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDBNAME")
		port     = os.Getenv("PGPORT")
		timeZone = os.Getenv("TIMEZONE")
	)

	dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	sslMode := "disable"
	if env == "production" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf(dsnFormat, host, user, password, dbname, port, sslMode, timeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db, nil

}
