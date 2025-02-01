package main

import (
    "log"
    "fmt"
    "net/http"

    "github.com/ingarondel/GO-APIDevelopment/config"
    "github.com/ingarondel/GO-APIDevelopment/internal/db"
    "github.com/ingarondel/GO-APIDevelopment/internal/handler"
)

func main() {
// TODO connection к db ты прокидываешь на уровень repository - там находятся все твои запросы к бд
// TODO repository ты прокидываешь на уровень service - это твоя основная бизнес логика и там у тебя в основном будут вызовы repository
// TODO service ты прокидываешь на уровень handler - там у тебя логика по валидации запросов, вызова методов service и формирование ответов
// TODO handler ты прокидываешь в route и там мапишь на endpoints

    cfg, err := config.LoadConfig()
    if err != nil {
      log.Fatal("Config loading failed:", err)
    }

    dbConnect, err := db.NewPostgresConnection()
    if err != nil {
      log.Fatal("Database connection failed:", err)
    }
    defer dbConnect.Close()

    if err := db.RunMigrations(dbConnect); err != nil {
      log.Fatal("Failed to run migrations:", err)
    }

    r := handler.Routes(dbConnect)

    serverAddress := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
    log.Printf("Server started on %s", serverAddress)
    if err := http.ListenAndServe(serverAddress, r); err != nil {
      log.Fatal("Server failed:", err)
    }
}
