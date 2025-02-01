package service

import (
    "context"
    "errors"

    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)

type CartItemService struct {
    repo *repository.CartItemRepository
}

func NewCartItemService(repo *repository.CartItemRepository) *CartItemService {
    return &CartItemService{repo: repo}
}

func (s *CartItemService) CreateCartItem(ctx context.Context, item *model.CartItem) error {
    if item.CartID <= 0 {
        return errors.New("invalid cart ID")
    }
    if item.Product == "" {
        return errors.New("product name cannot be empty")
    }
    if item.Quantity <= 0 {
        return errors.New("quantity must be greater than zero")
    }

    return s.repo.CreateCartItem(ctx, item)
}

func (s *CartItemService) GetCartItems(ctx context.Context, cartID int64) ([]model.CartItem, error) {
    if cartID <= 0 {
        return nil, errors.New("invalid cart ID")
    }
    return s.repo.GetCartItems(ctx, cartID)
}

func (s *CartItemService) DeleteCartItem(ctx context.Context, cartID, itemID int64) error {
    if cartID <= 0 {
        return errors.New("invalid cart ID")
    }
    if itemID <= 0 {
        return errors.New("invalid item ID")
    }
    return s.repo.DeleteCartItem(ctx, cartID, itemID)
}
