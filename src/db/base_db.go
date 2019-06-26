package base_db

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func Query(sql_str string) (*sql.Rows, *sql.Stmt, *sql.DB) {
	fmt.Println("query::sql_str = ", sql_str)
	db, _ := sql.Open("mysql", "zxlworking:working@tcp(zxltest.zicp.vip:42278)/joke")
	//db, _ := sql.Open("mysql", "zxlworking:working@tcp(103.46.128.20:42278)/joke")

	stmt, prepare_err := db.Prepare(sql_str)
	fmt.Println("query::prepare_err = ", prepare_err)
	rows, query_err := stmt.Query()
	fmt.Println("query::query_err = ", query_err)

	return rows, stmt, db
}
