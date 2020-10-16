package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"http_handle"
)

var (
	patterArr = []string{
		"/test",
		"/test/qsbk/list",
		"/test/qsbk/detail",
		"/test/qsbk/comment/list",
		"/test/star/evaluate_self",
		"/test/star/star_info_list",
		"/test/maoyan/movie_list",
		"/test/maoyan/movie_detail",
		"/test/upload_file",
		"/test/user/register"}

	handleArr = []func(http.ResponseWriter, *http.Request){
		http_handle.SayHelloName,
		http_handle.QsbkList,
		http_handle.QsbkDetail,
		http_handle.QsbkCommentList,
		http_handle.StarEvaluateSelf,
		http_handle.StarInfoList,
		http_handle.MaoYanMovieList,
		http_handle.MaoYanMovieDetail,
		http_handle.UploadFile,
		http_handle.UserRegister}
)

func main() {
	for index, patterUrl := range patterArr {
		http.HandleFunc(patterUrl, handleArr[index])
	}

	temp_path := "/home/mi/zxl/workspace/my_github/go_test_server/src/temp_file/"
	fmt.Println("temp_path = ", temp_path)

	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	http.Handle("/", http.FileServer(http.Dir(p)))

	err := http.ListenAndServe(":9090", nil) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
