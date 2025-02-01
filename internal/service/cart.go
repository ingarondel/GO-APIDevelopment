package service

import (
    "context"
    "errors"
    
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
        return model.Cart{}, err
    }
    return *cart, nil
}

func (s *CartService) GetCart(ctx context.Context, id int64) (model.Cart, error) {
    if id <= 0 {
        return model.Cart{}, errors.New("invalid cart ID")
    }
    return s.repo.GetCart(ctx, id)
}
