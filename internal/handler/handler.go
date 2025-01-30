package handler

import (
    "database/sql"
    "github.com/gorilla/mux"

    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)
// TODO сюда передаешь только уровень handler
// TODO     r := mux.NewRouter() создаешь тут и возвращаешь в main
func Routes(r *mux.Router, db *sql.DB) {
    // TODO инициализируется на уровне repository
        // т.е. передается как зависимость для структуры Repository
    cartItemRepo := repository.NewCartItemRepository(db)
    // TODO инициализируется на уровне handler
    // т.е. передается как зависимость для структуры Handler
    cartItemHandler := NewCartItemHandler(cartItemRepo)

    // TODO инициализируется на уровне repository
    // т.е. передается как зависимость для структуры Repository
    cartRepo := repository.NewCartRepository(db)
        // TODO инициализируется на уровне handler
    // т.е. передается как зависимость для структуры Handler
    cartHandler := NewCartHandler(cartRepo, cartItemRepo)

    r.HandleFunc("/carts", cartHandler.CreateCart).Methods("POST")
    r.HandleFunc("/carts/{cartId}", cartHandler.GetCart).Methods("GET")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.AddCartItem).Methods("POST")
    r.HandleFunc("/carts/{cartId}/items/{itemId}", cartItemHandler.DeleteCartItem).Methods("DELETE")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.GetCartItems).Methods("GET")
}
//внутри хендлера должен быть хендлер.го и маршруты
