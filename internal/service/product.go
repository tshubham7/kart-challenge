package service

import (
	"context"

	"github.com/tshubham7/kart-challenge/internal/service/validator"
	"github.com/tshubham7/kart-challenge/pkg/db"
	"github.com/tshubham7/kart-challenge/pkg/model"
)

type httpService struct {
	store      db.Store
	validation validator.OrderValidator
}

type HttpService interface {
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetProducts(ctx context.Context) ([]model.Product, error)
	CreateOrder(ctx context.Context, couponCode string, items []model.OrderItem) (*model.Order, error)
}

func NewHttpService(store db.Store, validation validator.OrderValidator) *httpService {
	return &httpService{store, validation}
}

func (s *httpService) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	return s.store.GetProductById(id)
}

func (s *httpService) GetProducts(ctx context.Context) ([]model.Product, error) {
	return s.store.GetProducts()
}
