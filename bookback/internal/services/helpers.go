package services

import "strconv"

// DeleteQuery returns a SQL query to delete a row from a table by its ID.
// e.g. DeleteQuery(table_name, id_name) -> DELETE FROM table_name WHERE id_name = $1
func DeleteQuery(tableName, idName string) string {
	sql := `DELETE FROM` + " " + tableName + ` WHERE ` + idName + ` = $1`

	return sql
}

// SelectWhere returns a SQL query to select all items from a table with a WHERE clause.
// e.g. SelectWhere(allItems, table_name, column1, column2) -> SELECT allItems... FROM table_name WHERE column1 = $1 AND column2 = $2
func SelectWhere(allItems func() string, tableName string, args ...string) string {
	sql := `SELECT ` + allItems() + ` FROM ` + tableName
	if len(args) > 0 {
		sql += ` WHERE ` + ParamsToQuery(" AND ", args...)
	}
	return sql
}

// ParamsToQuery returns a string of SQL query parameters.
// e.g. ParamsToQuery(", ", "column1", "column2") -> column1 = $1, column2 = $2
func ParamsToQuery(joiner string, args ...string) (sql string) {
	for i, arg := range args {
		sql += arg + ` = $` + strconv.Itoa(i+1)
		if i < len(args)-1 {
			sql += joiner
		}
	}
	return
}
