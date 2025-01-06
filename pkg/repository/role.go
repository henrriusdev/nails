package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/store"
)

type Role struct {
	Base[model.Role]
}

func NewRole(db store.Queryable) *Role {
	return &Role{
		Base: Base[model.Role]{
			Table: "role",
			DB:    db,
		},
	}
}
