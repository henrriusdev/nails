package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/supabase-community/postgrest-go"
)

type Appointment struct {
	Base[model.Appointment]
}

func NewAppointment(client *postgrest.Client) *Appointment {
	return &Appointment{
		Base: Base[model.Appointment]{
			Table:  "appointment",
			Client: client,
		},
	}
}
