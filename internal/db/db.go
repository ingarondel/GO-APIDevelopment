package db

import (
	"database/sql"
	"fmt"

	"github.com/ingarondel/GO-APIDevelopment/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func NewPostgresConnection() (*sql.DB, error) { 
    cfg, err := config.LoadConfig()
    if err != nil {
        return nil, err
    }

    connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	connect, err := sql.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }

    if err = connect.Ping(); err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %w", err)
    }

    return connect, nil
}

func RunMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "./internal/db//migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
