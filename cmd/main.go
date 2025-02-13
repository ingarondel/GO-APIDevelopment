package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ingarondel/GO-APIDevelopment/config"
	"github.com/ingarondel/GO-APIDevelopment/db"
	"github.com/ingarondel/GO-APIDevelopment/internal/handler"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config loading failed:", err)
	}

	dbConnect, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer dbConnect.Close()

	cartHandler := handler.NewCartHandler(dbConnect)
	cartItemHandler := handler.NewCartItemHandler(dbConnect)

	r := handler.Routes(cartHandler, cartItemHandler)

	serverAddress := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Server started on %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
