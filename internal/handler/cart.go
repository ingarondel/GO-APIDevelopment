package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)

type CartHandler struct {
    cartRepo *repository.CartRepository
    cartItemRepo  *repository.CartItemRepository
}

func NewCartHandler(cartRepo *repository.CartRepository, cartItemRepo *repository.CartItemRepository) *CartHandler {
    return &CartHandler{cartRepo, cartItemRepo}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
    cart := &model.Cart{}
    if err := h.cartRepo.CreateCart(cart); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)

    cart, err := h.cartRepo.GetCart(cartID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    items, err := h.cartItemRepo.GetCartItems(cartID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "id": cart.ID,
        "items": items,
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
