package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/store"
)

type Customer struct {
	Base[model.Customer]
}

func NewCustomer(db store.Queryable) *Customer {
	return &Customer{
		Base: Base[model.Customer]{
			Table: "customer",
			DB:    db,
		},
	}
}
