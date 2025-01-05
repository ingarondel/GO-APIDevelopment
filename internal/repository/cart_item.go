package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/ingarondel/GO-APIDevelopment/internal/model"
)

type CartItemRepository struct {
    db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
    return &CartItemRepository{db}
}

func (r *CartItemRepository) GetCartItems(ctx context.Context, cartID int64) ([]model.CartItem, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    query := "SELECT id, cart_id, product, quantity FROM cart_items WHERE cart_id = $1"
    var items []model.CartItem
    err := r.db.Select(&items, ctx, query, cartID)
    return items, err
}

func (r *CartItemRepository) CreateCartItem(ctx context.Context, item *model.CartItem) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    query := "INSERT INTO cart_items (cart_id, product, quantity) VALUES ($1, $2, $3) RETURNING id"
    return r.db.QueryRowContext(ctx, query, item.CartID, item.Product, item.Quantity).Scan(&item.ID)
}

func (r *CartItemRepository) DeleteCartItem(ctx context.Context, cartID, itemID int64) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    query := "DELETE FROM cart_items WHERE id = $1 AND cart_id = $2"
    _, err := r.db.Exec(ctx, query, itemID, cartID)
    return err
}
