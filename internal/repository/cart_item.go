package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
)

type CartItemRepository struct {
    db *sqlx.DB
}

func NewCartItemRepository(db *sqlx.DB) *CartItemRepository {
    return &CartItemRepository{db}
}

func (r *CartItemRepository) GetCartItems(cartID int64) ([]model.CartItem, error) {
    query := "SELECT id, cart_id, product, quantity FROM cart_items WHERE cart_id = $1"
    var items []model.CartItem
    err := r.db.Select(&items, query, cartID)
    return items, err
}

func (r *CartItemRepository) CreateCartItem(item *model.CartItem) error {
    query := "INSERT INTO cart_items (cart_id, product, quantity) VALUES ($1, $2, $3) RETURNING id"
    return r.db.QueryRow(query, item.CartID, item.Product, item.Quantity).Scan(&item.ID)
}

func (r *CartItemRepository) DeleteCartItem(cartID, itemID int64) error {
    query := "DELETE FROM cart_items WHERE id = $1 AND cart_id = $2"
    _, err := r.db.Exec(query, itemID, cartID)
    return err
}
