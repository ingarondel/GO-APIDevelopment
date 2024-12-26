package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/db"
    "github.com/ingarondel/GO-APIDevelopment/internal/route"
)

func main() {
    if err := db.Connect(); err != nil {
        log.Fatal("Database connection failed:", err)
    }

    r := mux.NewRouter()
    
    route.Routes(r)

    log.Println("Server started on :3000")
    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Fatal("Server failed:", err)
    }
}
