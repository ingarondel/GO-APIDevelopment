package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
    "github.com/ingarondel/GO-APIDevelopment/internal/route"
)

func main() {
    db, err := repository.Connect()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    defer db.Close()

    r := mux.NewRouter()
    route.Routes(r, db)

    log.Println("Server started on :3000")
    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Fatal("Server failed:", err)
    }
}
