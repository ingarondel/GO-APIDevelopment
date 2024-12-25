package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/ingarondel/GO-APIDevelopment/model" 
)

type CartRepository struct {
    db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
    return &CartRepository{db}
}

func (r *CartRepository) CreateCart(cart *models.Cart) error {
    query := "INSERT INTO carts DEFAULT VALUES RETURNING id"
    return r.db.QueryRowx(query).Scan(&cart.ID)
}
