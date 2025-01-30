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
type CartHandler struct {
    cartRepo     *repository.CartRepository
    cartItemRepo *repository.CartItemRepository
}

// TODO cюда передаются зависимости от уровня service
func NewCartHandler(cartRepo *repository.CartRepository, cartItemRepo *repository.CartItemRepository) *CartHandler {
    return &CartHandler{cartRepo, cartItemRepo}
}
// TODO необходимо создать контекст с таймаутом
func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
    // TODO CreateCart - по логике мне кажется ему не нужно передавать  &model.Cart{}, тк ему впринципе для создания cart
    // ничего не нужно передавать
    cart := &model.Cart{}

    ctx := r.Context()  
    // TODO CreateCart - должен возвращать созданную cart
    if err := h.cartRepo.CreateCart(ctx, cart); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cart)
}
// TODO необходимо создать контекст с таймаутом
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    // TODO нужно обязатель оборачивать ошибки 
    // if err!=nil{
    //     ...
    // }
    cartID, _ := strconv.ParseInt(vars["cartId"], 10, 64)


    ctx := r.Context()
    cart, err := h.cartRepo.GetCart(ctx, cartID)
// TODO необходимо обработать ошибку, когда cart не найдена(404) и также учесть что могут быть другие ошибки(500)
    if err != nil {
// TODO для возврата ошибок используй тот же    
//  w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(some_thing)
// но нужно возвращать свою структуру ошибки 
// type ErrorResponse struct{
// message 'json:message'
// }
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    items, err := h.cartItemRepo.GetCartItems(ctx, cartID)
    if err != nil {
        // TODO необходимо обработать ошибку, когда cart не найдена(404) и также учесть что могут быть другие ошибки(500)
        // TODO для возврата ошибок используй тот же    
//  w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(some_thing)
// но нужно возвращать свою структуру ошибки 
// type ErrorResponse struct{
// message 'json:message'
// }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // TODO создай на уровне model структуру GetCartReponce и опиши там поля необходимые для возврата
    response := map[string]interface{}{
        "id":    cart.ID,
        "items": items,
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
