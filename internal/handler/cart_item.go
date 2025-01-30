package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)
// TODO тут могут быть только зависимости от уровня service
type CartItemHandler struct {
    cartItemRepo *repository.CartItemRepository
}
// TODO cюда передаются зависимости от уровня service
func NewCartItemHandler(cartItemRepo *repository.CartItemRepository) *CartItemHandler {
    return &CartItemHandler{cartItemRepo}
}
// TODO необходимо создать контекст с таймаутом
func (h *CartItemHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)
     // TODO нужно обязатель оборачивать ошибки 
    // if err!=nil{
    //     ...
    // }
    // TODO для каждого запроса лучше создавать свою структуру в этом случае нужно model.AddCartItemRequest
    // тк мы хотим различать структуры для запросов и ответов и структуры для бд
    var item model.CartItem
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil || item.Product == "" || item.Quantity <= 0 {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    item.CartID = cartID


    ctx := r.Context()  
    if err := h.cartItemRepo.CreateCartItem(ctx, &item); err != nil {
        // TODO необходимо обработать ошибку, когда cart не найдена(404) и также учесть что могут быть другие ошибки(500)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}
// TODO необходимо создать контекст с таймаутом
func (h *CartItemHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)
    itemID, _ := strconv.ParseInt(vars["itemId"], 10, 64)
     // TODO нужно обязатель оборачивать ошибки 
    // if err!=nil{
    //     ...
    // }

    ctx := r.Context()  
    if err := h.cartItemRepo.DeleteCartItem(ctx, cartID, itemID); err != nil {
                // TODO необходимо обработать ошибку, когда cart не найдена(404) и также учесть что могут быть другие ошибки(500)
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
// TODO необходимо создать контекст с таймаутом
func (h *CartItemHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)
// TODO нужно обязатель оборачивать ошибки 
    // if err!=nil{
    //     ...
    // }
    ctx := r.Context()
    items, err := h.cartItemRepo.GetCartItems(ctx, cartID)
    if err != nil {
     // TODO необходимо обработать ошибку, когда cart не найдена(404) и также учесть что могут быть другие ошибки(500)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(items)
}
