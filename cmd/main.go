package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/db"
    "github.com/ingarondel/GO-APIDevelopment/internal/handler"
)

func main() {
    // TODO нужно создать папку config и там создать фалйик config.go
    // внутри этого файлика тебе нужно
    // 1 распарсить полностью свой env файл в структуру к примеру 
    // type Config struct {
//     ServerHost string
//     ServerPort string
//     PostgressHost string
//     PostgressPort string
//...
// }
// и соотвествнно вернуть этот конфиг

// TODO connection к db ты прокидываешь на уровень repository - там находятся все твои запросы к бд
// TODO repository ты прокидываешь на уровень service - это твоя основная бизнес логика и там у тебя в основном будут вызовы repository
// TODO service ты прокидываешь на уровень handler - там у тебя логика по валидации запросов, вызова методов service и формирование ответов
// TODO handler ты прокидываешь в route и там мапишь на endpoints



    dbConnect, err := db.NewPostgresConnection()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    defer dbConnect.Close()

    if err := db.RunMigrations(dbConnect); err != nil {
        log.Fatal("Failed to run migrations:", err)
    }

    r := mux.NewRouter()
    handler.Routes(r, dbConnect)

    log.Println("Server started on :3000")
    if err := http.ListenAndServe(":3000", r); err != nil { // TODO":3000" - нужно взять эти данные из config это твои ServerHost&ServerPort
        log.Fatal("Server failed:", err)
    }
}