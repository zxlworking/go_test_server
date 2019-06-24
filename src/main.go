package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func qsbk_hot_pic_list(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	db, _ := sql.Open("mysql", "zxlworking:working"+"@tcp(zxltest.zicp.vip:42278)/"+"joke?charset=utf8mb4")
	stmt, _ := db.Prepare("SELECT content FROM joke")
	rows, _ := stmt.Query("content")

	if rows.Next() {
		var content string
		rows.Scan(&content)
		fmt.Println("qsbk_hot_pic_list::content = ", qsbk_hot_pic_list)
	}
	defer rows.Close()
	defer stmt.Close()
	defer db.Close()
}

func main() {
	http.HandleFunc("/test", sayhelloName) //设置访问的路由
	http.HandleFunc("/test/qsbk_hot_pic/list", qsbk_hot_pic_list)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
