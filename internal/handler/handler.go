package handler

import (
    "database/sql"
    "github.com/gorilla/mux"

    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
    "github.com/ingarondel/GO-APIDevelopment/internal/service"
)

func Routes(db *sql.DB) *mux.Router {
    r := mux.NewRouter()

    cartRepo := repository.NewCartRepository(db)
    cartItemRepo := repository.NewCartItemRepository(db)

    cartService := service.NewCartService(cartRepo)
    cartItemService := service.NewCartItemService(cartItemRepo)

    cartHandler := NewCartHandler(cartService, cartItemService)
    cartItemHandler := NewCartItemHandler(cartItemService)

    r.HandleFunc("/carts", cartHandler.CreateCart).Methods("POST")
    r.HandleFunc("/carts/{cartId}", cartHandler.GetCart).Methods("GET")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.AddCartItem).Methods("POST")
    r.HandleFunc("/carts/{cartId}/items/{itemId}", cartItemHandler.DeleteCartItem).Methods("DELETE")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.GetCartItems).Methods("GET")

    return r
}
