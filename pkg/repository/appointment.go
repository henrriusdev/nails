package repository

import (
	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/store"
)

type Appointment struct {
	Base[model.Appointment]
}

func NewAppointment(db store.Queryable) *Appointment {
	return &Appointment{
		Base: Base[model.Appointment]{
			Table: "appointment",
			DB:    db,
		},
	}
}
