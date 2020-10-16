package http_handle

import (
	base_db "db"
	"fmt"
	"net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var simpleParseResult SimpleParseResult
	simpleParseResult.parse(r)

	userName, userNameErr := simpleParseResult.ParseResult["user_name"]
	passWord, passWordErr := simpleParseResult.ParseResult["pass_word"]
	if userNameErr || passWordErr {
		fmt.Println("UserRegister Err")
	} else {
		fmt.Println("userName", userName)
		fmt.Println("passWord", passWord)
	}

	var dbUtils base_db.DBUtils

	queryColumn := []string{
		"id", "author_nick_name", "author_gender", "author_age", "author_img_url", "content", "thumb_img_url", "stats_vote_content", "stats_comment_content", "stats_comment_detail_url", "md5",
	}
	queryResult := dbUtils.Query("joke", queryColumn, "", "id", "desc", 0, 10)

	for k, v := range queryResult {
		fmt.Println("k = ", k)
		for k2, v2 := range v {
			fmt.Print(" ")
			fmt.Print("k2 = ", k2)
			fmt.Print(" ")
			fmt.Print("v2 = ", v2)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
