package db

import "github.com/tshubham7/kart-challenge/pkg/model"

type Store interface {
	GetProductById(id string) (*model.Product, error)
	GetProducts() ([]model.Product, error)
	CreateProduct(item *model.Product) error
	CreateOrder(item *model.Order) error
}
