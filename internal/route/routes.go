package route

import (
    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/handler"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
    "github.com/ingarondel/GO-APIDevelopment/internal/db"
)

func Routes(r *mux.Router) {
    cartItemRepo := repository.NewCartItemRepository(db.DB)
    cartItemHandler := handler.NewCartItemHandler(cartItemRepo)

    cartRepo := repository.NewCartRepository(db.DB)
    cartHandler := handler.NewCartHandler(cartRepo, cartItemRepo)

    r.HandleFunc("/carts", cartHandler.CreateCart).Methods("POST")
    r.HandleFunc("/carts/{cartId}", cartHandler.GetCart).Methods("GET")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.AddCartItem).Methods("POST")
    r.HandleFunc("/carts/{cartId}/items/{itemId}", cartItemHandler.DeleteCartItem).Methods("DELETE")
    r.HandleFunc("/carts/{cartId}/items", cartItemHandler.GetCartItems).Methods("GET")
}
