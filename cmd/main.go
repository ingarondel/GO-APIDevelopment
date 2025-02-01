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
