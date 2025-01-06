package service

import (
	"context"

	"github.com/henrriusdev/nails/pkg/model"
	"github.com/henrriusdev/nails/pkg/repository"
)

type Appointment struct {
	repo *repository.Appointment
}

func NewAppointment(repo *repository.Appointment) *Appointment {
	return &Appointment{repo: repo}
}

func (s *Appointment) Create(ctx context.Context, a model.Appointment) (model.Appointment, error) {
	return s.repo.InsertOne(ctx, a)
}
