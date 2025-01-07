package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/supabase-community/postgrest-go"
)

type Product struct {
	Base[model.Product]
}

func NewProduct(client *postgrest.Client) *Product {
	return &Product{
		Base: Base[model.Product]{
			Table:  "product",
			Client: client,
		},
	}
}
