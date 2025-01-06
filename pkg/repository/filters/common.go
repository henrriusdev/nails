package filters

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/henrriusdev/nails/pkg/model"
)

func IdSelectFilter(id string, tableAlias ...string) SelectFilterBuilder {
	column := "id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, id)
}

func UserIdSelectFilter(userId string, tableAlias ...string) SelectFilterBuilder {
	column := "user_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, userId)
}

func AccountIdSelectFilter(id string, tableAlias ...string) SelectFilterBuilder {
	column := "account_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, id)
}

func LinkedAccountIdSelectFilter(id string, tableAlias ...string) SelectFilterBuilder {
	column := "linked_account_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, id)
}

func CycleIdSelectFilter(id string, tableAlias ...string) SelectFilterBuilder {
	column := "cycle_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, id)
}

func CategoryIdSelectFilter(categoryId string, tableAlias ...string) SelectFilterBuilder {
	column := "category_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, categoryId)
}

func MethodIdSelectFilter(id string, tableAlias ...string) SelectFilterBuilder {
	column := "method_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, id)
}

func MerchantIdSelectFilter(id string, tableAlias ...string) SelectFilterBuilder {
	column := "merchant_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, id)
}

func NameSelectFilter(name string, tableAlias ...string) SelectFilterBuilder {
	column := "name"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, name)
}

// EndCurrentDateSelectFilter creates a filter for comparing end dates
func EndCurrentDateSelectFilter(tableAlias ...string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		column := "end_date"
		if len(tableAlias) > 0 {
			column = tableAlias[0] + "." + column
		}
		return q.Where(goqu.I(column).Gte(goqu.L("CURRENT_DATE - INTERVAL '1 day'")))
	}
}

// TransactionTypeFilter handles the transaction type logic by reusing filters
func TransactionTypeSelectFilter(transactionType, column string, tableAlias ...string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		if len(tableAlias) > 0 {
			column = tableAlias[0] + "." + column
		}
		switch transactionType {
		case "income":
			return GenericColumnSelectFilter(column, "Income", "")(q)
		case "expense":
			return NumericComparisonSelectFilter(column, 0, model.ComparisonOperatorGreater)(q)
		default:
			return q
		}
	}
}

func StartDateSelectFilter(startDate string, tableAlias ...string) SelectFilterBuilder {
	column := "start_date"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, startDate)
}

func EndDateSelectFilter(endDate string, tableAlias ...string) SelectFilterBuilder {
	column := "end_date"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, endDate)
}

func MonthSelectFilter(month interface{}, tableAlias ...string) SelectFilterBuilder {
	column := "month"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, month)
}

func DateSelectFilter(startDate, endDate *time.Time, tableAlias ...string) SelectFilterBuilder {
	column := "date"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}

	var start, end *time.Time
	if startDate != nil {
		start = startDate
	}
	if endDate != nil {
		end = endDate
	}

	return TimeSelectFilter(start, end, column)
}

func CategoryTypeSelectFilter(categoryType string, tableAlias ...string) SelectFilterBuilder {
	column := "type"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsSelectFilter(column, categoryType)
}

func TransactionTypeUpdateFilter(transactionType, column string, tableAlias ...string) UpdateFilterBuilder {
	return func(q *goqu.UpdateDataset) *goqu.UpdateDataset {
		if len(tableAlias) > 0 {
			column = tableAlias[0] + "." + column
		}
		switch transactionType {
		case "income":
			return GenericColumnUpdateFilter(column, "Income", "")(q)
		case "expense":
			return NumericComparisonUpdateFilter(column, 0, string(model.ComparisonOperatorGreater))(q)
		default:
			return q
		}
	}
}

// ### Update filters ###
func UserIdUpdatetFilter(userId string, tableAlias ...string) UpdateFilterBuilder {
	column := "user_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, userId)
}

func CategoryIdUpdatetFilter(categoryId string, tableAlias ...string) UpdateFilterBuilder {
	column := "category_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, categoryId)
}

func LinkedAccountIdUpdatetFilter(id string, tableAlias ...string) UpdateFilterBuilder {
	column := "linked_account_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, id)
}

func MerchantIdUpdatetFilter(id string, tableAlias ...string) UpdateFilterBuilder {
	column := "merchant_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, id)
}

func NameUpdatetFilter(name string, tableAlias ...string) UpdateFilterBuilder {
	column := "name"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, name)
}

func IdUpdateFilter(id string, tableAlias ...string) UpdateFilterBuilder {
	column := "id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, id)
}

func AccountIdUpdateFilter(id string, tableAlias ...string) UpdateFilterBuilder {
	column := "account_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, id)
}

func CycleIdUpdateFilter(id string, tableAlias ...string) UpdateFilterBuilder {
	column := "cycle_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, id)
}

func MethodIdUpdateFilter(id string, tableAlias ...string) UpdateFilterBuilder {
	column := "method_id"
	if len(tableAlias) > 0 {
		column = tableAlias[0] + "." + column
	}
	return IsUpdateFilter(column, id)
}

func SubqueryFilter(subquery *goqu.SelectDataset, condition string) SelectFilterBuilder {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.Where(goqu.L(condition, subquery))
	}
}

func NotExistsFilter(subquery *goqu.SelectDataset) SelectFilterBuilder {
	return SubqueryFilter(subquery, "NOT EXISTS (?)")
}

func CountColumn(column string) exp.SQLFunctionExpression {
	return goqu.COUNT(goqu.I(column))
}

func CountDistinctColumn(column string) exp.SQLFunctionExpression {
	return goqu.COUNT(goqu.DISTINCT(goqu.I(column)))
}

func DateTruncMonth(column string) exp.LiteralExpression {
	return goqu.L("DATE_TRUNC('month', ?)", goqu.I(column))
}

func CountDistinctDateTruncMonth(column string) exp.SQLFunctionExpression {
	return goqu.COUNT(goqu.DISTINCT(DateTruncMonth(column)))
}

func MaxColumn(column string) exp.SQLFunctionExpression {
	return goqu.MAX(column)
}

func MinColumn(column string) exp.SQLFunctionExpression {
	return goqu.MIN(column)
}

func MaxDateIntervalFilter(column string, interval string) exp.LiteralExpression {
	return goqu.L("MAX(?) >= CURRENT_DATE - INTERVAL '"+interval+"'", goqu.I(column))
}

func MinDateIntervalFilter(column string, interval string) exp.LiteralExpression {
	return goqu.L("MIN(?) <= CURRENT_DATE - INTERVAL '"+interval+"'", goqu.I(column))
}
