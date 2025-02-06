package service

import (
    "context"
    "log"

    "github.com/ingarondel/GO-APIDevelopment/internal/model"
    "github.com/ingarondel/GO-APIDevelopment/internal/repository"
)

type CartService struct {
    repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
    return &CartService{repo: repo}
}

func (s *CartService) CreateCart(ctx context.Context) (model.Cart, error) {
    cart := &model.Cart{}
    err := s.repo.CreateCart(ctx, cart)
    if err != nil {
      log.Printf("Failed to create cart: %v", err)
      return model.Cart{}, err
    }
    return *cart, nil
}

func (s *CartService) GetCart(ctx context.Context, id int64) (model.Cart, error) {
    cart, err := s.repo.GetCart(ctx, id)
    if err != nil {
      log.Printf("Failed to get cart ID: %d: %v", id, err)
      return model.Cart{}, err
    }
    return cart, nil
}
