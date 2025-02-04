package service

import (
    "context"
    "log"

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
    err := s.repo.CreateCartItem(ctx, item)
    if err != nil {
      log.Printf("Failed to create cart item: %v", err)
      return err
    }
    return nil
}

func (s *CartItemService) GetCartItems(ctx context.Context, cartID int64) ([]model.CartItem, error) {
    items, err := s.repo.GetCartItems(ctx, cartID)
    if err != nil {
      log.Printf("Failed to get cart items for cart %d: %v", cartID, err)
      return nil, err
    }
    return items, nil
}

func (s *CartItemService) DeleteCartItem(ctx context.Context, cartID, itemID int64) error {
    err := s.repo.DeleteCartItem(ctx, cartID, itemID)
    if err != nil {
      log.Printf("Failed to delete cart item for cart: %d, item: %d; %v", cartID, itemID, err)
      return err
    }
    return nil
}
