package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "errors"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/service"
)

type CartItemHandler struct {
    cartItemService *service.CartItemService
}

func NewCartItemHandler(cartItemService *service.CartItemService) *CartItemHandler {
    return &CartItemHandler{
        cartItemService: cartItemService,
    }
}

func (h *CartItemHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    
    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
      if err!=nil{
        http.Error(w, "Invalid cart ID", http.StatusBadRequest)
        return
      }

    var cartitem model.CartItem
    if err := h.cartItemService.CreateCartItem(ctx, &cartItem); err != nil {
      http.Error(w, "Invalid input", http.StatusBadRequest)
      return
    }

    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    item := model.CartItem{
      Product:  cartitem.Product,
      Quantity: cartitem.Quantity,
      CartID:   cartID,
    }

    if err := h.cartItemService.CreateCartItem(ctx, &item); err != nil {
      if errors.Is(err, service.ErrCartNotFound) {
        http.Error(w, "Cart not found", http.StatusNotFound)
        return
      }
      http.Error(w, "Internal server error", http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func (h *CartItemHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
    if err != nil {
      http.Error(w, "Invalid cart ID", http.StatusBadRequest)
      return
    }

    itemID, err := strconv.ParseInt(vars["itemId"], 10, 64)
    if err != nil {
      http.Error(w, "Invalid item ID", http.StatusBadRequest)
      return
    }

    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    if err := h.cartItemService.DeleteCartItem(ctx, cartID, itemID); err != nil {
       if errors.Is(err, service.ErrCartNotFound) {
        http.Error(w, "Cart not found", http.StatusNotFound)
        return
      }
      http.Error(w, "Internal server error", http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (h *CartItemHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
    if err != nil {
      http.Error(w, "Invalid cart ID", http.StatusBadRequest)
      return
    }
    
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    items, err := h.cartItemService.GetCartItems(ctx, cartID)
    if err != nil {
       if errors.Is(err, service.ErrCartNotFound) {
        http.Error(w, "Cart not found", http.StatusNotFound)
        return
      }
      http.Error(w, "Internal server error", http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(items)
}
