package model

type (
	ComparisonOperator string
)

const (
	ComparisonOperatorGreater    ComparisonOperator = "gt"
	ComparisonOperatorLess       ComparisonOperator = "lt"
	ComparisonOperatorEqual      ComparisonOperator = "eq"
	ComparisonOperatorNotEqual   ComparisonOperator = "neq"
	ComparisonOperatorIs         ComparisonOperator = "is"
	ComparisonOperatorLike       ComparisonOperator = "like"
	ComparisonOperatorIn         ComparisonOperator = "in"
	ComparisonOperatorGte        ComparisonOperator = "gte"
	ComparisonOperatorLte        ComparisonOperator = "lte"
	ComparisonOperatorNotLike    ComparisonOperator = "notLike"
	ComparisonOperatorILike      ComparisonOperator = "iLike"
	ComparisonOperatorNotILike   ComparisonOperator = "notILike"
	ComparisonOperatorBetween    ComparisonOperator = "between"
	ComparisonOperatorNotBetween ComparisonOperator = "notBetween"
)
