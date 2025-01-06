package filters

import (
	"strings"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/henrriusdev/nails/pkg/model"
)

// SelectFilterBuilder defines a function signature for filters
type SelectFilterBuilder func(*goqu.SelectDataset) *goqu.SelectDataset

// ApplyFilters applies all filters to the query
func ApplyFilters(q *goqu.SelectDataset, filters ...SelectFilterBuilder) *goqu.SelectDataset {
	for _, filter := range filters {
		q = filter(q)
	}
	return q
}

// GenericColumnFilter handles filtering based on different types (string, bool, etc.)
func GenericColumnSelectFilter[T comparable](column string, value T, zeroValue T) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if value != zeroValue {
			return q.Where(goqu.I(column).Eq(value))
		}
		return q
	}
}

// TimeFilter filters based on start and/or end time
func TimeSelectFilter(startDate, endDate *time.Time, column string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if startDate != nil {
			q = q.Where(goqu.I(column).Gte(*startDate))
		}
		if endDate != nil {
			q = q.Where(goqu.I(column).Lte(*endDate))
		}
		return q
	}
}

// LimitOffsetFilter applies limit and offset for pagination
func LimitOffsetFilter(limit, offset uint) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if limit > 0 {
			q = q.Limit(limit)
		}
		if offset > 0 {
			q = q.Offset(offset)
		}
		return q
	}
}

// IsFilter handles filtering for various types, including NULL values
func IsSelectFilter(column string, value interface{}) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		switch v := value.(type) {
		case bool:
			return GenericColumnSelectFilter(column, v, false)(q)
		case nil:
			return q.Where(goqu.I(column).IsNull())
		case string:
			if v == "NOT NULL" {
				return q.Where(goqu.I(column).IsNotNull())
			} else {
				return GenericColumnSelectFilter(column, v, "")(q)
			}
		default:
			return q.Where(goqu.I(column).Eq(v))
		}
	}
}

func IsExSelectFilter(ex goqu.Ex) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(ex)
	}
}

// SQLIsFilter handles filtering for a column with a SQL expression
func IsSqlSelectFilter(column, expression string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("? = ?", goqu.I(column), goqu.L(expression)))
	}
}

// NotFilter negates the condition produced by another filter
func NotSelectFilter(filter SelectFilterBuilder) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		// Apply the filter to get its expression
		clauses := filter(goqu.From()).GetClauses()

		// Check if the filter has a WHERE clause to negate
		if clauses.Where() != nil && clauses.Where().Expression() != nil {
			// Negate the WHERE clause
			return q.Where(goqu.L("NOT ?", clauses.Where().Expression()))
		}

		// Return the dataset unchanged if no WHERE clause exists
		return q
	}
}

// OrFilter takes the conditions produced by another filter and append them to an or
func OrSelectFilter(filters ...SelectFilterBuilder) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		conditions := []goqu.Expression{}
		for _, filter := range filters {
			clauses := filter(goqu.From()).GetClauses()
			if clauses.Where() != nil && clauses.Where().Expression() != nil {
				conditions = append(conditions, clauses.Where().Expression())
			}
		}
		if len(conditions) > 0 {
			return q.Where(goqu.Or(conditions...))
		}
		return q
	}
}

// NumericComparisonFilter creates a filter for numeric comparisons
func NumericComparisonSelectFilter(column string, value interface{}, op model.ComparisonOperator) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		switch op {
		case model.ComparisonOperatorLess:
			return q.Where(goqu.I(column).Lt(value))
		case model.ComparisonOperatorGreater:
			return q.Where(goqu.I(column).Gt(value))
		case model.ComparisonOperatorEqual:
			return q.Where(goqu.I(column).Eq(value))
		default:
			return q
		}
	}
}

// InFilter applies an "IN" condition on the specified column for a list of values
func InSelectFilter(column string, values []string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if len(values) > 0 {
			return q.Where(goqu.I(column).In(values))
		}
		return q
	}
}

func InColumnSelectFilter(column string, value interface{}) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.Ex{
			column: goqu.Op{"@>": []interface{}{value}},
		})
	}
}

func CurrentDateSelectFilter(column string, op string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		switch op {
		case "lt":
			return q.Where(goqu.I(column).Lt(goqu.L("CURRENT_DATE")))
		case "gt":
			return q.Where(goqu.I(column).Gt(goqu.L("CURRENT_DATE")))
		default:
			return q
		}
	}
}

func ExtractYearSelectFilter(column string, year int) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("EXTRACT(YEAR FROM ?) = ?", goqu.I(column), year))
	}
}

func ExtractDayFilter(column, value string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("EXTRACT(DAY FROM ?) = ?", goqu.I(column), value))
	}
}

func ExtractDayIntervalFilter(column, value, interval string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("EXTRACT(DAY FROM ? - INTERVAL ?) = ?", goqu.I(column), interval, value))
	}
}

func ExtractDayEqualFilter(column string, value int) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("EXTRACT(DAY FROM ?) = ?", goqu.I(column), value))
	}
}

func DateLessEqualCurrentFilter(column string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("? <= CURRENT_DATE", goqu.I(column)))
	}
}

func DateLessEqualIntervalFilter(column, interval string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("? <= CURRENT_DATE - INTERVAL ?", goqu.I(column), interval))
	}
}

func DateGreaterEqualIntervalFilter(column, interval string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("? >= CURRENT_DATE - INTERVAL ?", goqu.I(column), interval))
	}
}

func DateIntervalFilter(column, interval string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("? >= CURRENT_DATE - INTERVAL ?", goqu.I(column), interval))
	}
}

func ModeNameFromSubQueryFilter(subQuery *goqu.SelectDataset) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L("t.name = (SELECT mode() WITHIN GROUP (ORDER BY name) FROM (?) as sub)", subQuery))
	}
}

func PayDayConditionsFilter(column string, payDay int) SelectFilterBuilder {
	return OrSelectFilter(
		ExtractDayEqualFilter(column, payDay),
		AndSelectFilter(
			ExtractDayEqualFilter("CURRENT_DATE - INTERVAL '1 day'", payDay),
			DateLessEqualIntervalFilter(column, "1 day"),
		),
		AndSelectFilter(
			ExtractDayEqualFilter("CURRENT_DATE", payDay),
			DateLessEqualCurrentFilter(column),
		),
	)
}

func AndSelectFilter(filters ...SelectFilterBuilder) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		for _, filter := range filters {
			q = filter(q)
		}
		return q
	}
}

// JoinSelectFilter creates a join filter
func JoinSelectFilter(table string, alias string, on goqu.Ex) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.LeftJoin(
			goqu.T(table).As(alias),
			goqu.On(on),
		)
	}
}

// IsLikeSelectFilter creates a filter for a LIKE condition
func IsLikeSelectFilter(column, value string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.I(column).Like("%" + value + "%"))
	}
}

func OnSelectFilter(columnPairs ...string) goqu.Ex {
	if len(columnPairs) == 0 || len(columnPairs)%2 != 0 {
		panic("Invalid number of column pairs")
	}

	ex := goqu.Ex{}
	for i := 0; i < len(columnPairs); i += 2 {
		ex[columnPairs[i]] = goqu.I(columnPairs[i+1])
	}
	return ex
}

func ExSelectFilter(columns []string, values []interface{}) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		for i := range columns {
			q = q.Where(goqu.Ex{
				columns[i]: values[i],
			})
		}
		return q
	}
}

func ExOpSelectFilter(column string, op string, value interface{}) goqu.Ex {
	return goqu.Ex{
		column: goqu.Op{op: value},
	}
}

// SelectColumnsFilter specifies which columns to select
func SelectColumnsFilter(columns ...interface{}) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		var columnNames []interface{}
		for _, column := range columns {
			columnNames = append(columnNames, goqu.I(column.(string)))
		}
		return q.Select(columnNames...)
	}
}

// SelectRawColumnsFilter specifies which columns to select using goqu.L
func SelectRawColumnsFilter(columns ...string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		var columnNames []interface{}
		for _, column := range columns {
			columnNames = append(columnNames, goqu.L(column))
		}
		return q.Select(columnNames...)
	}
}

// JsonbFieldSelectFilter creates a filter for a specific field in a JSONB column
func JsonbFieldIsSelectFilter(column, field, value string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.L(column + "->>'" + field + "'").Eq(value))
	}
}

// JsonbObjectFieldSelectFilter creates a filter for a specific field in a JSONB column that returns a JSON object
func JsonbObjectFieldIsSelectFilter(column, field, value string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.L(column + "->'" + field + "'").Eq(value))
	}
}

// JsonbFieldNotIsSelectFilter creates a filter for a specific field in a JSONB column
func JsonbFieldNotIsSelectFilter(column, field, value string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.L(column + "->>'" + field + "'").Neq(value))
	}
}

func JsonbObjectFieldNotIsSelectFilter(column, field, value string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.L(column + "->'" + field + "'").Neq(value))
	}
}

// OrderByFilter adds ordering to the query
func OrderByFilter(column string, direction string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if direction == "desc" {
			return q.Order(goqu.I(column).Desc())
		}
		return q.Order(goqu.I(column).Asc())
	}
}

// GroupByFilter adds grouping to the query
func GroupByFilter(columns ...string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			if strings.HasPrefix(col, "DATE_TRUNC") {
				cols[i] = goqu.L(col)
				continue
			}
			cols[i] = goqu.I(col)
		}
		return q.GroupBy(cols...)
	}
}

// HavingFilter adds a HAVING clause
func HavingFilter(exs ...exp.Expression) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Having(exs...)
	}
}

func OrderBySumFilter(column string, direction string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if direction == "desc" {
			return q.Order(goqu.SUM(column).Desc())
		}
		return q.Order(goqu.SUM(column).Asc())
	}
}

func CountFilter(column string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Select(goqu.COUNT(column))
	}
}

// UpdateFilterBuilder defines a function signature for filters
type UpdateFilterBuilder func(*goqu.UpdateDataset) *goqu.UpdateDataset

// ApplyUpdateFilters applies all filters to the update query
func ApplyUpdateFilters(q *goqu.UpdateDataset, filters ...UpdateFilterBuilder) *goqu.UpdateDataset {
	for _, filter := range filters {
		q = filter(q)
	}
	return q
}

// GenericColumnUpdateFilter handles filtering based on different types (string, bool, etc.)
func GenericColumnUpdateFilter[T comparable](column string, value T, zeroValue T) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		if value != zeroValue {
			return q.Where(goqu.I(column).Eq(value))
		}
		return q
	}
}

// TimeUpdateFilter filters based on start and/or end time
func TimeUpdateFilter(startDate, endDate *time.Time, column string) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		if startDate != nil {
			q = q.Where(goqu.I(column).Gte(*startDate))
		}
		if endDate != nil {
			q = q.Where(goqu.I(column).Lte(*endDate))
		}
		return q
	}
}

// IsUpdateFilter handles filtering for various types, including NULL values
func IsUpdateFilter(column string, value interface{}) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		switch v := value.(type) {
		case bool:
			return GenericColumnUpdateFilter(column, v, false)(q)
		case nil:
			return q.Where(goqu.I(column).IsNull())
		case string:
			if v == "NOT NULL" {
				return q.Where(goqu.I(column).IsNotNull())
			} else {
				return GenericColumnUpdateFilter(column, v, "")(q)
			}
		default:
			return q.Where(goqu.I(column).Eq(v))
		}
	}
}

// NotUpdateFilter negates the condition produced by another filter
func NotUpdateFilter(filter UpdateFilterBuilder) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		// Apply the filter to get its expression
		clauses := filter(q).GetClauses()

		// Check if the filter has a WHERE clause to negate
		if clauses.Where() != nil && clauses.Where().Expression() != nil {
			// Negate the WHERE clause
			return q.Where(goqu.L("NOT ?", clauses.Where().Expression()))
		}

		// Return the dataset unchanged if no WHERE clause exists
		return q
	}
}

// OrUpdateFilter takes the conditions produced by another filter and append them to an or
func OrUpdateFilter(filters ...UpdateFilterBuilder) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		conditions := []goqu.Expression{}
		for _, filter := range filters {
			clauses := filter(q).GetClauses()
			if clauses.Where() != nil && clauses.Where().Expression() != nil {
				conditions = append(conditions, clauses.Where().Expression())
			}
		}
		if len(conditions) > 0 {
			return q.Where(goqu.Or(conditions...))
		}
		return q
	}
}

// NumericComparisonUpdateFilter creates a filter for numeric comparisons
func NumericComparisonUpdateFilter(column string, value int, op string) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		switch op {
		case "lt":
			return q.Where(goqu.I(column).Lt(value))
		case "gt":
			return q.Where(goqu.I(column).Gt(value))
		default:
			return q
		}
	}
}

// InUpdateFilter applies an "IN" condition on the specified column for a list of values
func InUpdateFilter(column string, values []string) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		if len(values) > 0 {
			return q.Where(goqu.I(column).In(values))
		}
		return q
	}
}

func CurrentDateUpdateFilter(column string, op string) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		switch op {
		case "lt":
			return q.Where(goqu.I(column).Lt(goqu.L("CURRENT_DATE")))
		case "gt":
			return q.Where(goqu.I(column).Gt(goqu.L("CURRENT_DATE")))
		default:
			return q
		}
	}
}

func UpdateOnConflictFilter(columns ...string) goqu.Record {
	updates := make(goqu.Record, len(columns))
	for _, column := range columns {
		updates[column] = goqu.L("EXCLUDED." + column)
	}
	return updates
}
