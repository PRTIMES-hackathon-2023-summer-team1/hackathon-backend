package config

import (
	"os"
)

type appInfo struct {
	PostgresInfo *PostgresInfo
}

type PostgresInfo struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

func LoadConfig() *appInfo {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	appInfo := &appInfo{
		PostgresInfo: &PostgresInfo{
			User:     dbUser,
			Password: dbPass,
			Database: dbName,
			Host:     dbHost,
			Port:     dbPort,
		},
	}
	return appInfo
}
