package model

type Cart struct {
	ID int64 `db:"id"`
	Items []CartItem `json:"items"`
}
