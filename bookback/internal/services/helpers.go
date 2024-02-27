package services

import "strconv"

func SelectWhere(allItems func() string, tableName string, args ...string) string {
	sql := `SELECT ` + allItems() + ` FROM ` + tableName
	if len(args) > 0 {
		sql += ` WHERE ` + ParamsToQuery(" AND ", args...)
	}
	return sql
}

func ParamsToQuery(joiner string, args ...string) (sql string) {
	for i, arg := range args {
		sql += arg + ` = $` + strconv.Itoa(i+1)
		if i < len(args)-1 {
			sql += joiner
		}
	}
	return
}
