package service

import (
	"context"

	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/repository"
)

type Product struct {
	repo *repository.Product
}

func NewProduct(repo *repository.Product) *Product {
	return &Product{repo: repo}
}

func (s *Product) Create(ctx context.Context, p model.Product) (model.Product, error) {
	return s.repo.InsertOne(ctx, p, "", "", false)
}
