package db

import (
	"database/sql"
	"fmt"
	"embed"

	"github.com/ingarondel/GO-APIDevelopment/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) { 
    connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.PostgresHost, cfg.User, cfg.Password, cfg.DBName, cfg.PostgresPort, cfg.SSLMode)

	connect, err := sql.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }

    if err = connect.Ping(); err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %w", err)
    }

    if err = runMigrations(connect); err != nil {
      return nil, err
    }

    return connect, nil
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	goose.SetBaseFS(migrationFiles)

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
