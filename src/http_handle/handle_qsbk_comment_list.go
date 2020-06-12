package http_handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"

	"data"
	base_db "db"
)

func QsbkCommentList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	var last_id = -1
	if len(r.Form["last_id"]) > 0 {
		last_id, _ = strconv.Atoi(r.Form["last_id"][0])
	}
	joke_id := r.Form["joke_id"][0]
	page, _ := strconv.Atoi(r.Form["page"][0])
	page_size, _ := strconv.Atoi(r.Form["page_size"][0])
	comment_type, _ := strconv.Atoi(r.Form["comment_type"][0])
	fmt.Println("last_id", last_id)
	fmt.Println("joke_id", joke_id)
	fmt.Println("page", page)
	fmt.Println("page_size", page_size)
	fmt.Println("comment_type", comment_type)

	var start_index = page * page_size
	var end_index = page_size
	fmt.Println("start_index", start_index)
	fmt.Println("end_index", end_index)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	var sql_where = " WHERE joke_id = " + joke_id
	if last_id > 0 {
		start_index = 0
		end_index = 10
		sql_where = sql_where + " AND id < " + strconv.Itoa(last_id)
		if comment_type > 0 {
			sql_where = sql_where + " AND comment_type > 0 "
		} else {
			sql_where = sql_where + " AND comment_type = 0 "
		}
	} else {
		if comment_type > 0 {
			sql_where = sql_where + " AND comment_type > 0 "
		} else {
			sql_where = sql_where + " AND comment_type = 0 "
		}
	}

	fmt.Println("sql_where", sql_where)

	var sql_str = "SELECT " +
		"id, joke_id,article_id,comment_user_id,comment_user_img,comment_user_nick_name,comment_user_gender,comment_user_age,comment_user_content,comment_type " +
		"FROM joke_comment " + sql_where +
		"ORDER BY id DESC LIMIT " + strconv.Itoa(start_index) + "," + strconv.Itoa(end_index)

	rows, stmt, db := base_db.Query(sql_str)

	var qsbkCommentList data.QsbkCommentList
	for rows.Next() {
		var id, joke_id, article_id, comment_user_id, comment_user_img, comment_user_nick_name, comment_user_gender, comment_user_age, comment_user_content, comment_type string
		rows.Scan(&id, &joke_id, &article_id, &comment_user_id, &comment_user_img, &comment_user_nick_name, &comment_user_gender, &comment_user_age, &comment_user_content, &comment_type)
		fmt.Println("query::result = ", id, joke_id, article_id, comment_user_id, comment_user_img, comment_user_nick_name, comment_user_gender, comment_user_age, comment_user_content, comment_type)

		var item = data.QsbkCommentItem{id, joke_id, article_id, comment_user_id, comment_user_img, comment_user_nick_name, comment_user_gender, comment_user_age, comment_user_content, comment_type}
		qsbkCommentList.ItemList = append(qsbkCommentList.ItemList, item)
	}
	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	responseResult, responseError := json.Marshal(qsbkCommentList)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
