package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ingarondel/GO-APIDevelopment/internal/errorsx"
	"github.com/ingarondel/GO-APIDevelopment/internal/model"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db}
}

func (r *CartRepository) CreateCart(ctx context.Context) (model.Cart, error) {
	cart := &model.Cart{}
	query := "INSERT INTO carts DEFAULT VALUES RETURNING id"

	err := r.db.QueryRowContext(ctx, query).Scan(&cart.ID)
	if err != nil {
		return model.Cart{}, fmt.Errorf("failed to create cart: %w", err)
	}
	return *cart, nil
}

func (r *CartRepository) GetCart(ctx context.Context, id int64) (model.Cart, error) {
	query := "SELECT id FROM carts WHERE id = $1"
	var cart model.Cart
	// псомотреть как решить проблему с 2м походом
	err := r.db.QueryRowContext(ctx, query, id).Scan(&cart.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cart, errorsx.ErrCartNotFound
		}
		return cart, fmt.Errorf("failed to get cart: %w", err)
	}
	return cart, nil
}
