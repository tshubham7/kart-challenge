package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tshubham7/kart-challenge/pkg/constants"
	"github.com/tshubham7/kart-challenge/pkg/model"
)

func (s *httpService) CreateOrder(ctx context.Context, couponCode string, items []model.OrderItem) (*model.Order, error) {
	if err := s.validation.ValidateCouponCode(ctx, couponCode); err != nil {
		return nil, err
	}

	products, err := s.fetchAllProducts(ctx, items)
	if err != nil {
		return nil, err
	}

	order := &model.Order{
		Id:             uuid.New().String(),
		Items:          items,
		Products:       products,
		CouponDiscount: s.getDiscountInPercentage(couponCode),
	}
	if err := s.store.CreateOrder(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *httpService) fetchAllProducts(_ context.Context, items []model.OrderItem) ([]model.Product, error) {
	products := make([]model.Product, len(items))
	for i, item := range items {
		prod, err := s.store.GetProductById(item.ProductId)
		if err != nil {
			return nil, errors.New(constants.ErrorCodeBadRequest)
		}
		products[i] = *prod
	}
	return products, nil
}

func (s *httpService) getDiscountInPercentage(couponCode string) int {
	// dummy implementation
	if couponCode == "" {
		return 0
	}
	return 15
}
