package repository

import (
	"os"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// TODO - не знаю зачем это,но думаю оно тебе не нужно
var DB *sql.DB
// TODO переименуй это как NewPostgressConnetion
// TODO перенеси это в db package
// TODO os.Getenv должно происходить в config.go сюда же ты только передаешь структуру Config


func Connect() (*sql.DB, error) { 
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// TODO для формирования строк исопльзуй fmt.Sprintf
	// connect -> connectionString
	connect := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
// TODO "postgres" - тоже достаешь и Config
// TODO database -> conn
database, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	return database, nil
}
