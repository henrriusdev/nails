package service

import (
	"context"

	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/repository"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	return &User{repo: repo}
}

func (s *User) Create(ctx context.Context, u model.User) (model.User, error) {
	return s.repo.InsertOne(ctx, u, "", "", false)
}
