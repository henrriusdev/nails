package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/store"
)

type User struct {
	Base[model.User]
}

func NewUser(db store.Queryable) *User {
	return &User{
		Base: Base[model.User]{
			Table: "users",
			DB:    db,
		},
	}
}
