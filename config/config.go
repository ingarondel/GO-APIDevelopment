package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    Host     string
    User     string
    Password string
    DBName   string
    Port     string
    SSLMode  string
}

func LoadConfig() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        return nil, err
    }

    return &Config{
        Host:     os.Getenv("DB_HOST"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        Port:     os.Getenv("DB_PORT"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }, nil
}
