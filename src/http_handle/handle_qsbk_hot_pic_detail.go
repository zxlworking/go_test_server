package http_handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"

	"data"
	base_db "db"
)

func QsbkHotPicDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	hot_pic_id := r.Form["hot_pic_id"][0]
	fmt.Println("hot_pic_id", hot_pic_id)

	var sql_str = "SELECT " +
		"id, hot_pic_id,article_id,stats_time,content,thumb_img_url " +
		"FROM hot_pic_joke_detail " +
		"WHERE hot_pic_id = " + hot_pic_id

	rows, stmt, db := base_db.Query(sql_str)

	var qsbkHotPicDetailInfo data.QsbkHotPicDetailInfo
	for rows.Next() {
		var id, hot_pic_id, article_id, stats_time, content, thumb_img_url string
		rows.Scan(&id, &hot_pic_id, &article_id, &stats_time, &content, &thumb_img_url)
		fmt.Println("query::result = ", id, hot_pic_id, article_id, stats_time, content, thumb_img_url)

		qsbkHotPicDetailInfo = data.QsbkHotPicDetailInfo{id, hot_pic_id, article_id, stats_time, content, thumb_img_url}
	}
	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	responseResult, responseError := json.Marshal(qsbkHotPicDetailInfo)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
