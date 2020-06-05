package base_db

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func Query(sql_str string) (*sql.Rows, *sql.Stmt, *sql.DB) {
	fmt.Println("query::sql_str = ", sql_str)
	db, open_db_err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/life")
	// db, open_db_err := sql.Open("mysql", "@tcp(zxltest.zicp.vip:42278)/life")
	//db, open_db_err := sql.Open("mysql", ":@tcp(103.46.128.49:42278)/life")
	fmt.Println("query::open_db_err = ", open_db_err)
	stmt, prepare_err := db.Prepare(sql_str)
	fmt.Println("query::prepare_err = ", prepare_err)
	rows, query_err := stmt.Query()
	fmt.Println("query::query_err = ", query_err)

	return rows, stmt, db
}
