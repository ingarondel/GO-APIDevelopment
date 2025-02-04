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
    "github.com/ingarondel/GO-APIDevelopment/internal/errorsx"
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
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    vars := mux.Vars(r)
    
    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
      if err!=nil || cartID <= 0{
        respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
        return
      }

    var cartitem model.CartItem
    if err := json.NewDecoder(r.Body).Decode(&cartitem); err != nil || cartitem.Product == "" || cartitem.Quantity <= 0 {
      respondWithError(w, http.StatusBadRequest, "Invalid input")
      return
    }

    item := model.CartItem{
      Product:  cartitem.Product,
      Quantity: cartitem.Quantity,
      CartID:   cartID,
    }

    if err := h.cartItemService.CreateCartItem(ctx, &item); err != nil {
      if errors.Is(err, errorsx.ErrCartNotFound) {
        respondWithError(w, http.StatusNotFound, "Cart not found")
        return
      }
      respondWithError(w, http.StatusInternalServerError, "Something went wrong")
      return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func (h *CartItemHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    vars := mux.Vars(r)

    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
    if err != nil || cartID <= 0{
      respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
      return
    }

    itemID, err := strconv.ParseInt(vars["itemId"], 10, 64)
    if err != nil || itemID <= 0 {
      respondWithError(w, http.StatusBadRequest, "Invalid item ID")
      return
    }

    if err := h.cartItemService.DeleteCartItem(ctx, cartID, itemID); err != nil {
       if errors.Is(err, errorsx.ErrCartNotFound) {
        respondWithError(w, http.StatusNotFound, "Cart not found")
        return
      }
      respondWithError(w, http.StatusInternalServerError, "Something went wrong")
      return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (h *CartItemHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    vars := mux.Vars(r)

    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
    if err != nil || cartID <= 0 {
      respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
      return
    }

    items, err := h.cartItemService.GetCartItems(ctx, cartID)
    if err != nil {
       if errors.Is(err, errorsx.ErrCartNotFound) {
        respondWithError(w, http.StatusNotFound, "Cart not found")
        return
      }
      respondWithError(w, http.StatusInternalServerError, "Something went wrong")
      return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(items)
}
