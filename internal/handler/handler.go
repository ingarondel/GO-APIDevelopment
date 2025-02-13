package handler

import (
	"github.com/gorilla/mux"
)

func Routes(cartHandler *CartHandler, cartItemHandler *CartItemHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/carts", cartHandler.CreateCart).Methods("POST")
	r.HandleFunc("/carts/{cartId}", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/carts/{cartId}/items", cartItemHandler.AddCartItem).Methods("POST")
	r.HandleFunc("/carts/{cartId}/items/{itemId}", cartItemHandler.DeleteCartItem).Methods("DELETE")
	r.HandleFunc("/carts/{cartId}/items", cartItemHandler.GetCartItems).Methods("GET")

	return r
}
