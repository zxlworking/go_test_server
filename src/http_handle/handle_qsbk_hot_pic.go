package http_handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"

	data "data"
	base_db "db"
)

func QsbkHotPicList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	page, _ := strconv.Atoi(r.Form["page"][0])
	page_size, _ := strconv.Atoi(r.Form["page_size"][0])
	fmt.Println("page", page)
	fmt.Println("page_size", page_size)

	var start_index = page * page_size
	var end_index = (page + 1) * page_size
	fmt.Println("start_index", start_index)
	fmt.Println("end_index", end_index)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	var sql_str = "SELECT " +
		"id, author_nick_name,author_gender,author_age,author_img_url,content,thumb_img_url,stats_vote_content,stats_comment_content,stats_comment_detail_url,md5 " +
		"FROM joke " +
		"ORDER BY id DESC LIMIT " + strconv.Itoa(start_index) + "," + strconv.Itoa(end_index)

	rows, stmt, db := base_db.Query(sql_str)

	var qsbkHotPicItemList data.QsbkHotPicItemList
	for rows.Next() {
		var id, author_nick_name, author_gender, author_age, author_img_url, content, thumb_img_url, stats_vote_content, stats_comment_content, stats_comment_detail_url, md5 string
		rows.Scan(&id, &author_nick_name, &author_gender, &author_age, &author_img_url, &content, &thumb_img_url, &stats_vote_content, &stats_comment_content, &stats_comment_detail_url, &md5)
		fmt.Println("query::result = ", id, author_nick_name, author_gender, author_age, author_img_url, content, thumb_img_url, stats_vote_content, stats_comment_content, stats_comment_detail_url, md5)

		var item = data.QsbkHotPicItem{id, author_nick_name, author_gender, author_age, author_img_url, content, thumb_img_url, stats_vote_content, stats_comment_content, stats_comment_detail_url, md5}
		qsbkHotPicItemList.ItemList = append(qsbkHotPicItemList.ItemList, item)
	}
	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	responseResult, responseError := json.Marshal(qsbkHotPicItemList)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
