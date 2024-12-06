package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tshubham7/kart-challenge/pkg/model"
)

func TestUnit_MemoryStore(t *testing.T) {
	store := newMemoryStore(false)

	product := &model.Product{
		Id:       "1",
		Name:     "Chicken Waffle",
		Price:    40,
		Category: "Waffle",
	}
	err := store.CreateProduct(product)
	assert.Nil(t, err)

	actualProduct, err := store.GetProductById(product.Id)
	assert.Nil(t, err)
	assert.Equal(t, product, actualProduct)

	expectedProducts := []model.Product{*product}
	actualProducts, err := store.GetProducts()
	assert.Nil(t, err)
	assert.Equal(t, expectedProducts, actualProducts)

	order := model.Order{
		Id:       uuid.New().String(),
		Products: actualProducts,
		Items: []model.OrderItem{
			{
				ProductId: product.Id,
				Quantity:  2,
			},
		},
	}
	err = store.CreateOrder(&order)
	assert.Nil(t, err)
}
