package handler

import (
    "database/sql"
    "github.com/gorilla/mux"

    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)

func Routes(r *mux.Router, db *sql.DB) {
    cartItemRepo := repository.NewCartItemRepository(db)
    cartItemHandler := NewCartItemHandler(cartItemRepo)

    cartRepo := repository.NewCartRepository(db)
    cartHandler := NewCartHandler(cartRepo, cartItemRepo)

    r.HandleFunc("/carts", cartHandler.CreateCart).Methods("POST")
    r.HandleFunc("/carts/{cartId}", cartHandler.GetCart).Methods("GET")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.AddCartItem).Methods("POST")
    r.HandleFunc("/carts/{cartId}/items/{itemId}", cartItemHandler.DeleteCartItem).Methods("DELETE")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.GetCartItems).Methods("GET")
}
//внутри хендлера должен быть хендлер.го и маршруты
