package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/store"
)

type Product struct {
	Base[model.Product]
}

func NewProduct(db store.Queryable) *Product {
	return &Product{
		Base: Base[model.Product]{
			Table: "product",
			DB:    db,
		},
	}
}
