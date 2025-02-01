package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "strconv"
    "time"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/service"
)

type ErrorResponse struct {
    Message string `json:"message"`
}

type CartHandler struct {
    cartService     *service.CartService
    cartItemService *service.CartItemService
}

func NewCartHandler(cartService *service.CartService, cartItemService *service.CartItemService) *CartHandler {
    return &CartHandler{
      cartService:     cartService,
      cartItemService: cartItemService,
    }
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    cart, err := h.cartService.CreateCart(ctx)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, "Failed to create cart")
      return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, err := strconv.ParseInt(vars["cartId"], 10, 64)
    if err!=nil{
      respondWithError(w, http.StatusBadRequest, "Invalid cart ID")
      return
    }
    
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    cart, err := h.cartService.GetCart(ctx, cartID)
    if err != nil {
      respondWithErrorFromMap(w, err)
      return
    }

   items, err := h.cartItemService.GetCartItems(ctx, cartID)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, "Failed to retrieve cart items")
      return
    }

    response := model.Cart{
      ID:    cart.ID,
      Items: items,
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

func respondWithErrorFromMap(w http.ResponseWriter, err error) {
    errorResponses := map[error]struct {
      code    int
      message string
    }{
      service.ErrInvalidCartID:  {http.StatusBadRequest, "Invalid cart ID"},
      service.ErrCartNotFound:    {http.StatusNotFound, "Cart not found"},
    }

    if response, ok := errorResponses[err]; ok {
      respondWithError(w, response.code, response.message)
    } else {
      respondWithError(w, http.StatusInternalServerError, "Failed to retrieve cart")
    }
}
