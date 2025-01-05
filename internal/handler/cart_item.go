package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)

type CartItemHandler struct {
    cartItemRepo *repository.CartItemRepository
}

func NewCartItemHandler(cartItemRepo *repository.CartItemRepository) *CartItemHandler {
    return &CartItemHandler{cartItemRepo}
}

func (h *CartItemHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)
    var item model.CartItem
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil || item.Product == "" || item.Quantity <= 0 {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    item.CartID = cartID

    ctx := r.Context()  
    if err := h.cartItemRepo.CreateCartItem(ctx, &item); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func (h *CartItemHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)
    itemID, _ := strconv.ParseInt(vars["itemId"], 10, 64)

    ctx := r.Context()  

    if err := h.cartItemRepo.DeleteCartItem(ctx, cartID, itemID); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *CartItemHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)
    ctx := r.Context()

    items, err := h.cartItemRepo.GetCartItems(ctx, cartID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(items)
}
