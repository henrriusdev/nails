package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/supabase-community/postgrest-go"
)

type Role struct {
	Base[model.Role]
}

func NewRole(client *postgrest.Client) *Role {
	return &Role{
		Base: Base[model.Role]{
			Table:  "role",
			Client: client,
		},
	}
}
