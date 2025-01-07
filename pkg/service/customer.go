package service

import (
	"context"

	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/repository"
)

type Customer struct {
	repo *repository.Customer
}

func NewCustomer(repo *repository.Customer) *Customer {
	return &Customer{repo: repo}
}

func (s *Customer) Create(ctx context.Context, c model.Customer) (model.Customer, error) {
	return s.repo.InsertOne(ctx, c, "", "", false)
}
