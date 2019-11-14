package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"http_handle"
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

func main() {
	http.HandleFunc("/test", sayhelloName) //设置访问的路由
	http.HandleFunc("/test/qsbk_hot_pic/list", http_handle.QsbkHotPicList)
	http.HandleFunc("/test/qsbk_hot_pic/detail", http_handle.QsbkHotPicDetail)
	http.HandleFunc("/test/star/evaluate_self", http_handle.StarEvaluateSelf)
	http.HandleFunc("/test/star/star_info_list", http_handle.StarInfoList)
	http.HandleFunc("/test/maoyan/movie_list", http_handle.MaoYanMovieList)
	http.HandleFunc("/test/maoyan/movie_detail", http_handle.MaoYanMovieDetail)

	temp_path, _ := os.Getwd()
	temp_path = temp_path + "/src/temp_file"
	fmt.Println("temp_path = ", temp_path)

	//p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//p, _ := filepath.Abs(filepath.Dir("/Users/zxl/Documents/workspace/go_test_server/src/temp_file"))
	p, _ := filepath.Abs(filepath.Dir(temp_path))
	http.Handle("/", http.FileServer(http.Dir(p)))

	err := http.ListenAndServe(":9090", nil) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
