package db

import (
	"fmt"

	"github.com/tshubham7/kart-challenge/pkg/model"
)

func seedStore(st Store) {
	seedProducts(st)
}

func seedProducts(st Store) {
	for i := 1; i < 11; i++ {
		product := &model.Product{
			Id:       fmt.Sprintf("%d", i),
			Name:     fmt.Sprintf("Product %d", i),
			Category: "Starter",
			Price:    (10 * i) + i,
		}
		st.CreateProduct(product)
	}
}
