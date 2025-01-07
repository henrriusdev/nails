package repository

import (
	"context"
	"errors"

	"github.com/supabase-community/postgrest-go"
)

type Repositories struct {
	Appointment *Appointment
	Customer    *Customer
	Product     *Product
	Role        *Role
	User        *User
}

var (
	ErrNotFound = errors.New("not found")
	ErrExists   = errors.New("already exists")
)

type Base[T any] struct {
	Client *postgrest.Client
	Table  string
}

func (b *Base[T]) GetAll(ctx context.Context, columns, count string, head bool) ([]T, error) {
	var results []T
	query := b.Client.From(b.Table).Select(columns, count, head)

	_, err := query.ExecuteTo(&results)
	if err != nil {
		return *new([]T), err
	}

	return results, nil
}

func (b *Base[T]) GetOne(ctx context.Context, columns, count string, head bool) (T, error) {
	query := b.Client.From(b.Table).Select(columns, count, head).Single()

	var result []T
	_, err := query.ExecuteTo(&result)
	if err != nil {
		return *new(T), err
	}

	return result[0], nil
}

func (b *Base[T]) GetOneById(ctx context.Context, id string, columns, count string, head bool) (T, error) {
	query := b.Client.From(b.Table).Select(columns, count, head).Eq("id", id).Single()

	var result []T
	_, err := query.ExecuteTo(&result)
	if err != nil {
		return *new(T), err
	}

	return result[0], nil
}

func (b *Base[T]) UpdateOneById(ctx context.Context, id, returning, count string, model T) (T, error) {
	query := b.Client.From(b.Table).Update(model, returning, count).Eq("id", id).Single()

	var updated []T
	_, err := query.ExecuteTo(&updated)
	if err != nil {
		return *new(T), err
	}

	return updated[0], nil
}

func (b *Base[T]) UpdateOne(ctx context.Context, model T, returning, count string) (T, error) {
	query := b.Client.From(b.Table).Update(model, returning, count).Single()

	var updated []T
	_, err := query.ExecuteTo(&updated)
	if err != nil {
		return *new(T), err
	}

	return updated[0], nil
}

func (b *Base[T]) InsertOne(ctx context.Context, model T, returning, count string, upsert bool, onConflict ...string) (T, error) {
	var query *postgrest.FilterBuilder
	if upsert {
		query = b.Client.From(b.Table).Insert(model, true, onConflict[0], returning, count)
	} else {
		query = b.Client.From(b.Table).Insert(model, false, "", returning, count)
	}

	var inserted []T
	_, err := query.ExecuteTo(&inserted)
	if err != nil {
		return *new(T), err
	}

	return inserted[0], nil
}

