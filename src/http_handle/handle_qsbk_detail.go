package http_handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"

	"data"
	base_db "db"
)

func QsbkDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	joke_id := r.Form["joke_id"][0]
	fmt.Println("joke_id", joke_id)

	var sql_str = "SELECT " +
		"id, joke_id,article_id,stats_time,content,thumb_img_url " +
		"FROM joke_detail " +
		"WHERE joke_id = " + joke_id

	rows, stmt, db := base_db.Query(sql_str)

	var qsbkHotPicDetailInfo data.QsbkHotPicDetailInfo
	for rows.Next() {
		var id, joke_id, article_id, stats_time, content, thumb_img_url string
		rows.Scan(&id, &joke_id, &article_id, &stats_time, &content, &thumb_img_url)
		fmt.Println("query::result = ", id, joke_id, article_id, stats_time, content, thumb_img_url)

		qsbkHotPicDetailInfo = data.QsbkHotPicDetailInfo{id, joke_id, article_id, stats_time, content, thumb_img_url}
	}
	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	responseResult, responseError := json.Marshal(qsbkHotPicDetailInfo)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
