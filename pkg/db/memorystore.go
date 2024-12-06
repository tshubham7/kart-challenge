package db

import (
	"errors"

	"github.com/tshubham7/kart-challenge/pkg/constants"
	"github.com/tshubham7/kart-challenge/pkg/model"
)

type memoryStore struct {
	products map[string]model.Product
	orders   map[string]model.Order
}

func newMemoryStore(seed bool) *memoryStore {
	store := &memoryStore{
		products: make(map[string]model.Product),
		orders:   make(map[string]model.Order),
	}
	if seed {
		seedStore(store)
	}
	return store
}

func (s *memoryStore) GetProducts() ([]model.Product, error) {
	products := make([]model.Product, len(s.products))
	i := 0
	for _, v := range s.products {
		products[i] = v
		i++
	}

	return products, nil
}

func (s *memoryStore) GetProductById(id string) (*model.Product, error) {
	product, ok := s.products[id]
	if !ok {
		return nil, errors.New(constants.ErrorCodeEntityNotFound)
	}
	return &product, nil
}

func (s *memoryStore) CreateProduct(item *model.Product) error {
	if _, ok := s.products[item.Id]; ok {
		return errors.New(constants.ErrorCodeEntityAlreadyExist)
	}

	s.products[item.Id] = *item
	return nil
}

func (s *memoryStore) CreateOrder(item *model.Order) error {
	s.orders[item.Id] = *item
	return nil
}
