package http_handle

import (
	"data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	base_db "db"
)

func MaoYanMovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MaoYanMovieList::start")

	r.ParseForm()
	movie_type, _ := strconv.Atoi(r.Form["movie_type"][0])
	page, _ := strconv.Atoi(r.Form["page"][0])
	page_size, _ := strconv.Atoi(r.Form["page_size"][0])
	fmt.Println("movie_type", movie_type)
	fmt.Println("page", page)
	fmt.Println("page_size", page_size)

	var start_index = page * page_size
	fmt.Println("start_index", start_index)

	table_name := "mao_yan_now"
	switch movie_type {
	case 1:
		table_name = "mao_yan_now"
		break
	case 2:
		table_name = "mao_yan_future"
		break
	case 3:
		table_name = "mao_yan_history"
		break
	}

	var sql_str = "SELECT " +
		" id, movie_id, movie_title, movie_poster_url, movie_detail_url, movie_type " +
		" FROM " + table_name +
		" LIMIT " + strconv.Itoa(start_index) + "," + strconv.Itoa(page_size)

	rows, stmt, db := base_db.Query(sql_str)

	var movieInfoListResponseBean data.MovieInfoListResponseBean

	for rows.Next() {
		var id, movie_id, movie_title, movie_poster_url, movie_detail_url string
		var movie_type int
		rows.Scan(&id, &movie_id, &movie_title, &movie_poster_url, &movie_detail_url, &movie_type)
		fmt.Println("query::result = ", id, movie_id, movie_title, movie_poster_url, movie_detail_url, movie_type)

		movieInfo := data.MovieInfo{id, movie_id, movie_title, movie_poster_url, movie_detail_url, movie_type}

		movieInfoListResponseBean.MovieInfoList = append(movieInfoListResponseBean.MovieInfoList, movieInfo)
	}

	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	if len(movieInfoListResponseBean.MovieInfoList) > 0 {
		movieInfoListResponseBean.BaseBean.Code = 0
		movieInfoListResponseBean.BaseBean.Desc = "success"
	} else {
		movieInfoListResponseBean.BaseBean.Code = 1
		movieInfoListResponseBean.BaseBean.Desc = "fail"
		var movieInfo data.MovieInfo
		movieInfoListResponseBean.MovieInfoList = append(movieInfoListResponseBean.MovieInfoList, movieInfo)
	}

	responseResult, responseError := json.Marshal(movieInfoListResponseBean)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
