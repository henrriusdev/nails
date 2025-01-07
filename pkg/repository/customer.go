package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/supabase-community/postgrest-go"
)

type Customer struct {
	Base[model.Customer]
}

func NewCustomer(client *postgrest.Client) *Customer {
	return &Customer{
		Base: Base[model.Customer]{
			Table:  "customer",
			Client: client,
		},
	}
}
