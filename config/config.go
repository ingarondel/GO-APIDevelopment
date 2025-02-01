package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    PostgresHost     string
    PostgresPort     string
    ServerHost       string
    ServerPort       string
    User             string
    Password         string
    DBName           string
    SSLMode          string
}

func LoadConfig() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        return nil, err
    }

    return &Config{
        PostgresHost:     os.Getenv("DB_HOST"),
        PostgresPort:     os.Getenv("DB_PORT"),
        ServerHost:       os.Getenv("SERVER_HOST"),
        ServerPort:       os.Getenv("SERVER_PORT"),  
        User:             os.Getenv("DB_USER"),
        Password:         os.Getenv("DB_PASSWORD"),
        DBName:           os.Getenv("DB_NAME"),
        SSLMode:          os.Getenv("DB_SSLMODE"),
    }, nil
}
