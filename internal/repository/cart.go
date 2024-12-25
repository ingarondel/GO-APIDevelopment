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

func (r *CartRepository) CreateCart(cart *model.Cart) error {
    query := "INSERT INTO carts DEFAULT VALUES RETURNING id"
    return r.db.QueryRowx(query).Scan(&cart.ID)
}

func (r *CartRepository) GetCart(id int64) (model.Cart, error) {
    query := "SELECT id FROM carts WHERE id = $1"
    var cart models.Cart
    err := r.db.Get(&cart, query, id)
    return cart, err
}
