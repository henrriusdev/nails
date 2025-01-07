package service

import (
	"context"

	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/repository"
)

type Role struct {
	repo *repository.Role
}

func NewRole(repo *repository.Role) *Role {
	return &Role{repo: repo}
}

func (s *Role) Create(ctx context.Context, r model.Role) (model.Role, error) {
	return s.repo.InsertOne(ctx, r, "", "", false)
}
