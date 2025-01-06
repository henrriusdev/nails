package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"

	"github.com/henrriusdev/nails/pkg/repository/filters"
	"github.com/henrriusdev/nails/pkg/store"
)

type Repositories struct {
	Appointment *Appointment
	Customer    *Customer
	Product     *Product
	Role        *Role
	User        *User
}

const (
	USER_TABLE_NAME              = "user"
	ACCOUNTS_TABLE_NAME          = "linked_accounts"
	MERCHANT_TABLE_NAME          = "merchants"
	MERCHANT_ANALYSIS_TABLE_NAME = "merchant_analysis"
	MERCHANT_SUBSCRIPTION_TABLE  = "merchant_subscriptions"
	MERCHANT_USER_RULE_TABLE     = "merchant_user_rules"

	USER_KYC_VERIFICATION_TABLE = "user_kyc_verifications"

	TRANSACTION_SYNC_TABLE = "transaction_sync_status"

	BUDGET_CUSTOMIZATION_STEP_TABLE = "budget_customization_steps"
	BUDGET_CATEGORY_TABLE           = "budget_categories"
)

const (
	ID_COLUMN                 = "id"
	NAME_COLUMN               = "name"
	TYPE_COLUMN               = "type"
	AMOUNT_COLUMN             = "amount"
	DATE_COLUMN               = "date"
	PENDING_COLUMN            = "pending"
	MERCHANT_ID_COLUMN        = "merchant_id"
	MERCHANT_ENTITY_ID_COLUMN = "merchant_entity_id"
	CATEGORY_ID_COLUMN        = "category_id"
	USER_ID_COLUMN            = "user_id"
	LOGO_URL_COLUMN           = "logo_url"
	LAST_MONTH_AMOUNT_COLUMN  = "last_month_amount"
	PERCENTAGE_COLUMN         = "percentage"
	KEPT_COLUMN               = "kept"
	STATUS_COLUMN             = "status"
	PAY_DAY_COLUMN            = "pay_day"
	LAST_TRANSACTION_COLUMN   = "last_transaction"
	FIRST_TRANSACTION_COLUMN  = "first_transaction"
	CURRENT_AMOUNT_COLUMN     = "current_amount"
)

const (
	MERCHANT_ANALYSIS_CONFLICT     = "user_id, merchant_id, month"
	MERCHANT_CONFLICT              = "user_id, merchant_id"
	BUDGET_CATEGORY_CONFLICT       = "user_id, category_id"
	TRIMMING_SUBSCRIPTION_CONFLICT = "user_id, merchant_id, subscription_id"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExists   = errors.New("already exists")
	dialect     = goqu.Dialect("postgres")
)

type Base[T any] struct {
	Store store.Queryable
	DB    store.Queryable
	Table string
}

func (b *Base[T]) MustBegin() store.Queryable {
	db := b.Store.(*sqlx.DB)
	b.DB = db
	t := db.MustBegin()
	b.Store = t
	return t
}

func (b *Base[T]) Rollback() {
	t := b.Store.(*sqlx.Tx)
	t.Rollback()
	b.Reset()
}

func (b *Base[T]) Commit() error {
	t := b.Store.(*sqlx.Tx)
	err := t.Commit()
	if err != nil {
		return err
	}
	return err
}

func (b *Base[T]) SetTx(t store.Queryable) {
	b.DB = b.Store
	b.Store = t
}

func (b *Base[T]) Reset(repos ...store.Transactable) {
	b.Store = b.DB
	for _, v := range repos {
		v.Reset()
	}
}

// Helper function to generate a query builder with base filters (soft delete)
func (b *Base[T]) baseQuery(aliases ...string) *goqu.SelectDataset {
	alias := ""
	if len(aliases) > 0 {
		alias = aliases[0] // Use the first one
	}
	var tableExp exp.Expression
	table := goqu.T(b.Table)

	if alias != "" {
		tableExp = table.As(alias) // Apply alias if present
	} else {
		tableExp = table // Use table directly without alias
	}

	softDeleteColumn := b.Table + ".deleted_at"
	if alias != "" {
		softDeleteColumn = alias + ".deleted_at"
	}
	return dialect.From(tableExp).Where(goqu.I(softDeleteColumn).IsNull())
}

func (b *Base[T]) BaseQueryUpdate() *goqu.UpdateDataset {
	return dialect.Update(b.Table)
}

func (b *Base[T]) BaseQueryInsert() *goqu.InsertDataset {
	return dialect.Insert(b.Table)
}

func (b *Base[T]) GetAll(ctx context.Context, queryFilters ...filters.SelectFilterBuilder) ([]T, error) {
	var results []T
	query := filters.ApplyFilters(b.baseQuery(), queryFilters...)
	q, args, err := query.ToSQL()
	if err != nil {
		return results, err
	}

	if err := b.Store.SelectContext(ctx, &results, q, args...); err != nil {
		return results, err
	}
	return results, nil
}

func (b *Base[T]) GetOne(ctx context.Context, filter ...filters.SelectFilterBuilder) (T, error) {
	query := filters.ApplyFilters(
		b.baseQuery(),
		filter...,
	)

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var result T
	if err := b.Store.GetContext(ctx, &result, q, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return *new(T), ErrNotFound
		}
		return *new(T), err
	}

	return result, nil
}

func (b *Base[T]) GetOneById(ctx context.Context, id string) (T, error) {
	query := filters.ApplyFilters(
		b.baseQuery(),
		filters.IsSelectFilter("id", id),
	)

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var result T
	if err := b.Store.GetContext(ctx, &result, q, args...); err != nil {
		return *new(T), err
	}

	return result, nil
}

func (b *Base[T]) UpdateOneById(ctx context.Context, id string, model T, updateFilters ...filters.UpdateFilterBuilder) (T, error) {
	query := b.BaseQueryUpdate().
		Set(model).
		Where(goqu.Ex{"id": id})

	query = filters.ApplyUpdateFilters(query, updateFilters...)
	query = query.Returning("*")

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var updated T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&updated); err != nil {
		return *new(T), err
	}

	return updated, nil
}

func (b *Base[T]) UpdateOne(ctx context.Context, model T, updateFilters ...filters.UpdateFilterBuilder) (T, error) {
	query := b.BaseQueryUpdate().Set(model)

	query = filters.ApplyUpdateFilters(query, updateFilters...)
	query = query.Returning("*")

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var updated T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&updated); err != nil {
		return *new(T), err
	}

	return updated, nil
}

func (b *Base[T]) UpsertOneDoUpdate(ctx context.Context, model T, updateFields goqu.Record, conflictColumns ...string) (T, error) {
	query := b.BaseQueryInsert().
		Rows(model).
		OnConflict(goqu.DoUpdate(conflictColumns[0], updateFields)).
		Returning("*")

	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var upserted T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&upserted); err != nil {
		return *new(T), err
	}

	return upserted, nil
}

func (b *Base[T]) UpsertOneDoNothing(ctx context.Context, model T, conflictColumns ...string) (T, error) {
	query := b.BaseQueryInsert().Rows(model)
	if len(conflictColumns) > 0 {
		query = query.OnConflict(
			goqu.DoNothing(),
		)
	}

	query = query.Returning("*")
	q, args, err := query.ToSQL()
	if err != nil {
		return *new(T), err
	}

	var upserted T
	if err := b.Store.QueryRowxContext(ctx, q, args...).StructScan(&upserted); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return *new(T), ErrNotFound
		}

		return *new(T), err
	}

	return upserted, nil
}

func (b *Base[T]) Update(ctx context.Context, model T, updateFilters ...filters.UpdateFilterBuilder) error {
	query := b.BaseQueryUpdate().Set(model)
	query = filters.ApplyUpdateFilters(query, updateFilters...)

	q, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	if _, err := b.Store.ExecContext(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (b *Base[T]) InsertOne(ctx context.Context, model T, filters ...filters.SelectFilterBuilder) (T, error) {
	insert := b.BaseQueryInsert().Rows(model).Returning("*")
	sql, args, err := insert.ToSQL()
	if err != nil {
		return *new(T), err
	}

	err = b.Store.QueryRowxContext(ctx, sql, args...).StructScan(&model)
	if err != nil {
		return *new(T), err
	}
	return model, nil
}

func (b *Base[T]) InsertMany(ctx context.Context, models []T, filters ...filters.SelectFilterBuilder) ([]T, error) {
	if len(models) == 0 {
		return *new([]T), nil
	}

	insert := b.BaseQueryInsert().Rows(models).Returning("*")
	sql, args, err := insert.ToSQL()
	if err != nil {
		return *new([]T), err
	}
	rows, err := b.Store.QueryxContext(ctx, sql, args...)
	if err != nil {
		return *new([]T), err
	}

	var results []T
	for rows.Next() {
		var m T
		err = rows.StructScan(&m)
		if err != nil {
			return *new([]T), err
		}
		results = append(results, m)
	}
	return results, nil
}
