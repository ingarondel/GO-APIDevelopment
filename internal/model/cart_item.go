package model

type CartItem struct {
	ID       int64  `db:"id"`
	CartID   int64  `db:"cart_id"`
	Product  string `db:"product"`
	Quantity int64  `db:"quantity"`
}
