package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
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

// TODO
// Создаешь папку db и там же создаешь папку migrations(внутрь нее кидаешь все свои миграции)
// в папке db нужно создать файлик db.go - там ты формируешь connection к бд и также поднимаешь свои миграции

// TODO connection к db ты прокидываешь на уровень repository - там находятся все твои запросы к бд
// TODO repository ты прокидываешь на уровень service - это твоя основная бизнес логика и там у тебя в основном будут вызовы repository
// TODO service ты прокидываешь на уровень handler - там у тебя логика по валидации запросов, вызова методов service и формирование ответов
// TODO handler ты прокидываешь в route и там мапишь на endpoints



    db, err := repository.Connect()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    defer db.Close()

    r := mux.NewRouter()
    handler.Routes(r, db)

    log.Println("Server started on :3000")
    if err := http.ListenAndServe(":3000", r); err != nil { // TODO":3000" - нужно взять эти данные из config это твои ServerHost&ServerPort
        log.Fatal("Server failed:", err)
    }
}