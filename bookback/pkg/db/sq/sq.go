package sq

import (
	"github.com/Masterminds/squirrel"
)

var SQ = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Eq = squirrel.Eq
type NotEq = squirrel.NotEq
type Or = squirrel.Or
type And = squirrel.And
type SelectBuilder = squirrel.SelectBuilder

func Insert(table string) squirrel.InsertBuilder {
	return SQ.Insert(table)
}

func Select(columns ...string) squirrel.SelectBuilder {
	return SQ.Select(columns...)
}

func Update(table string) squirrel.UpdateBuilder {
	return SQ.Update(table)
}
