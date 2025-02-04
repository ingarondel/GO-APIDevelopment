package errorsx

import "errors"

var (
	ErrInvalidCartID = errors.New("invalid cart ID")
    ErrCartNotFound  = errors.New("cart not found")
)
