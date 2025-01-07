package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/supabase-community/postgrest-go"
)

type User struct {
	Base[model.User]
}

func NewUser(client *postgrest.Client) *User {
	return &User{
		Base: Base[model.User]{
			Table:  "users",
			Client: client,
		},
	}
}
