package base_db

import (
	"fmt"
	"strconv"
	"strings"
)

type DBUtils struct {
}

func (dbUtils *DBUtils) Query(tableName string, queryColumn []string,
	whereStr string,
	orderColumn string,
	descStr string,
	limitStartIndex int,
	limitEndIndex int) (queryResult map[string][]string) {
	fmt.Println("whereStr = ", whereStr)

	var querySqlStr = "*"
	for index, queryItem := range queryColumn {
		if index == 0 {
			querySqlStr = queryItem + ", "
		} else if index == len(queryColumn)-1 {
			querySqlStr = querySqlStr + queryItem + " "
		} else {
			querySqlStr = querySqlStr + queryItem + ", "
		}
	}
	fmt.Println("querySqlStr = ", querySqlStr)

	var sqlStr = "SELECT " + querySqlStr + " FROM " + tableName
	if strings.Count(whereStr, "")-1 > 0 {
		sqlStr = sqlStr + " WHERE " + whereStr
	}
	sqlStr = sqlStr + " ORDER BY " + orderColumn + " " + descStr + " " +
		" LIMIT " + strconv.Itoa(limitStartIndex) + "," + strconv.Itoa(limitEndIndex)

	fmt.Println("sqlStr = ", sqlStr)

	rows, stmt, db := Query(sqlStr)

	queryResult = make(map[string][]string)
	index := 0
	for rows.Next() {
		fmt.Println("rows.Next = ", index)
		var queryDest []string
		// for _, queryItem := range queryColumn {
		// 	queryDest = append(queryDest, &queryItem)
		// }

		rows.Scan(queryDest)

		for _, destItem := range queryDest {
			fmt.Println("destItem = ", *destItem)
			//queryResult[strconv.Itoa(index)] = append(queryResult[strconv.Itoa(index)], *destItem)
		}

		index = index + 1
	}
	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	return queryResult
}
