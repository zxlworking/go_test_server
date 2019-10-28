package http_handle

import (
	"data"
	"encoding/json"
	"fmt"
	"net/http"

	base_db "db"
)

func MaoYanMovieDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MaoYanMovieDetail::start")

	r.ParseForm()
	movie_id := r.Form["movie_id"][0]
	fmt.Println("movie_id", movie_id)

	table_name := "mao_yan_detail"

	var sql_str = "SELECT " +
		" id, movie_id, movie_avatar_url, movie_name, movie_en_name," +
		" movie_category, movie_duration, movie_release_info, movie_release_time," +
		" movie_release_area, movie_score_content," +
		" movie_stats_people_count_content,movie_stats_people_count_unit_content," +
		" movie_box_value_content, movie_box_unit_content, introduce_content" +
		" FROM " + table_name +
		" WHERE movie_id = " + movie_id

	rows, stmt, db := base_db.Query(sql_str)

	var movieDetailInfoListResponseBean data.MovieDetailInfoListResponseBean

	for rows.Next() {
		var id, movie_id, movie_avatar_url, movie_name, movie_en_name string
		var movie_category, movie_duration, movie_release_info, movie_release_time string
		var movie_release_area, movie_score_content string
		var movie_stats_people_count_content, movie_stats_people_count_unit_content string
		var movie_box_value_content, movie_box_unit_content, introduce_content string

		rows.Scan(
			&id, &movie_id, &movie_avatar_url, &movie_name, &movie_en_name,
			&movie_category, &movie_duration, &movie_release_info, &movie_release_time,
			&movie_release_area, &movie_score_content,
			&movie_stats_people_count_content, &movie_stats_people_count_unit_content,
			&movie_box_value_content, &movie_box_unit_content, &introduce_content,
		)

		fmt.Println("query::result = ", id, movie_id, movie_avatar_url, movie_name, movie_en_name,
			movie_category, movie_duration, movie_release_info, movie_release_time,
			movie_release_area, movie_score_content,
			movie_stats_people_count_content, movie_stats_people_count_unit_content,
			movie_box_value_content, movie_box_unit_content, introduce_content)

		movieDetailInfo := data.MovieDetailInfo{id, movie_id, movie_avatar_url, movie_name, movie_en_name,
			movie_category, movie_duration, movie_release_info, movie_release_time,
			movie_release_area, movie_score_content,
			movie_stats_people_count_content, movie_stats_people_count_unit_content,
			movie_box_value_content, movie_box_unit_content, introduce_content}

		movieDetailInfoListResponseBean.MovieDetailInfoList = append(movieDetailInfoListResponseBean.MovieDetailInfoList, movieDetailInfo)
	}

	defer rows.Close()
	defer stmt.Close()
	defer db.Close()

	if len(movieDetailInfoListResponseBean.MovieDetailInfoList) > 0 {
		movieDetailInfoListResponseBean.BaseBean.Code = 0
		movieDetailInfoListResponseBean.BaseBean.Desc = "success"
	} else {
		movieDetailInfoListResponseBean.BaseBean.Code = 1
		movieDetailInfoListResponseBean.BaseBean.Desc = "fail"
		var movieDetailInfo data.MovieDetailInfo
		movieDetailInfoListResponseBean.MovieDetailInfoList = append(movieDetailInfoListResponseBean.MovieDetailInfoList, movieDetailInfo)
	}

	responseResult, responseError := json.Marshal(movieDetailInfoListResponseBean)
	fmt.Println("query::responseError = ", responseError)
	fmt.Println("query::responseResult = ", string(responseResult))
	fmt.Fprint(w, string(responseResult))
}
