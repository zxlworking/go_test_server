package http_handle

import (
	"data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	base_db "db"
)

func StarInfoList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StarInfoList::start")
	r.ParseForm()
	page, _ := strconv.Atoi(r.Form["page"][0])
	page_size, _ := strconv.Atoi(r.Form["page_size"][0])
	fmt.Println("page", page)
	fmt.Println("page_size", page_size)
	var start_index = page * page_size
	fmt.Println("start_index", start_index)

	var sql_str = "SELECT " +
		" id, star_name,star_img_url,star_detail_url,face_id " +
		" FROM star_info " +
		" LIMIT " + strconv.Itoa(start_index) + "," + strconv.Itoa(page_size)

	rows, stmt, db := base_db.Query(sql_str)

	var evaluateSelfResponseBean data.EvaluateSelfResponseBean

	for rows.Next() {
		var id, star_name, star_img_url, star_detail_url, face_id string
		rows.Scan(&id, &star_name, &star_img_url, &star_detail_url, &face_id)
		fmt.Println("query::result = ", id, star_name, star_img_url, star_detail_url, face_id)

		var starInfoItem data.StarInfo
		var evaluateSelfInfoItem data.EvaluateSelfInfo

		starInfoItem = data.StarInfo{id, star_name, star_img_url, star_detail_url, face_id}
		evaluateSelfInfoItem.StarInfo = starInfoItem

		evaluateSelfResponseBean.EvaluateSelfInfoList = append(evaluateSelfResponseBean.EvaluateSelfInfoList, evaluateSelfInfoItem)
	}

	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	if len(evaluateSelfResponseBean.EvaluateSelfInfoList) > 0 {
		evaluateSelfResponseBean.BaseBean.Code = 0
		evaluateSelfResponseBean.BaseBean.Desc = "success"
	} else {
		evaluateSelfResponseBean.BaseBean.Code = 1
		evaluateSelfResponseBean.BaseBean.Desc = "fail"
		var starInfoItem data.StarInfo
		var evaluateSelfInfoItem data.EvaluateSelfInfo
		evaluateSelfInfoItem.StarInfo = starInfoItem
		evaluateSelfResponseBean.EvaluateSelfInfoList = append(evaluateSelfResponseBean.EvaluateSelfInfoList, evaluateSelfInfoItem)
	}

	responseResult, responseError := json.Marshal(evaluateSelfResponseBean)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
