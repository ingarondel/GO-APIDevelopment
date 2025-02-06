package repository

import (
    "context"
    "database/sql"
    "errors"
    "fmt"

    "github.com/ingarondel/GO-APIDevelopment/internal/errorsx"
    "github.com/ingarondel/GO-APIDevelopment/internal/model"
)

type CartItemRepository struct {
    db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
    return &CartItemRepository{db}
}

func (r *CartItemRepository) GetCartItems(ctx context.Context, cartID int64) ([]model.CartItem, error) {
    query := "SELECT id, cart_id, product, quantity FROM cart_items WHERE cart_id = $1"
    rows, err := r.db.QueryContext(ctx, query, cartID)
    if err != nil {
        return nil, fmt.Errorf("failed to get cart item: %w", err)
    }
    defer rows.Close()

    var items []model.CartItem

    for rows.Next() {
      var item model.CartItem
      if err := rows.Scan(&item.ID, &item.CartID, &item.Product, &item.Quantity); err != nil {
        return nil, fmt.Errorf("failed to scan cart item: %w", err)
      }
      items = append(items, item)
    }

    return items, nil
}

func (r *CartItemRepository) CreateCartItem(ctx context.Context, item *model.CartItem) error {
    var cartExists bool
    checkCartQuery := "SELECT EXISTS(SELECT 1 FROM carts WHERE id = $1)"

    err := r.db.QueryRowContext(ctx, checkCartQuery, item.CartID).Scan(&cartExists)
    if err != nil {
        return fmt.Errorf("failed to find cart: %w", err)
    }
    if !cartExists {
        return errorsx.ErrCartNotFound
    }

    query := "INSERT INTO cart_items (cart_id, product, quantity) VALUES ($1, $2, $3) RETURNING id"
    err = r.db.QueryRowContext(ctx, query, item.CartID, item.Product, item.Quantity).Scan(&item.ID)
    if err != nil {
      if errors.Is(err, sql.ErrNoRows) {
        return fmt.Errorf("cart with ID %d not found", item.CartID)
      }
        return fmt.Errorf("failed to create cart item: %w", err)
    }

    return nil
}

func (r *CartItemRepository) DeleteCartItem(ctx context.Context, cartID, itemID int64) error {
    var cartExists bool
    checkCartQuery := "SELECT EXISTS(SELECT 1 FROM carts WHERE id = $1)"
    err := r.db.QueryRowContext(ctx, checkCartQuery, cartID).Scan(&cartExists)
    if err != nil {
        return fmt.Errorf("failed to find cart: %w", err)
    }
    if !cartExists {
        return errorsx.ErrCartNotFound
    }

    var itemExists bool
    checkItemQuery := "SELECT EXISTS(SELECT 1 FROM cart_items WHERE id = $1 AND cart_id = $2)"
    err = r.db.QueryRowContext(ctx, checkItemQuery, itemID, cartID).Scan(&itemExists)
    if err != nil {
        return fmt.Errorf("failed to check if cart item exists: %w", err)
    }
    if !itemExists {
        return errorsx.ErrCartItemNotFound
    }


    query := "DELETE FROM cart_items WHERE id = $1 AND cart_id = $2"

    _, err = r.db.ExecContext(ctx, query, itemID, cartID)
    if err != nil {
        return fmt.Errorf("failed to delete cart item: %w", err)
    }

    return nil
}
