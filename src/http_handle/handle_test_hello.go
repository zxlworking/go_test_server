package http_handle

import (
	"fmt"
	"net/http"
	"strings"
)

func SayHelloName(w http.ResponseWriter, r *http.Request) {
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

	// 获取请求报文的内容长度
	len := r.ContentLength
	// 新建一个字节切片，长度与请求报文的内容长度相同
	body := make([]byte, len)
	// 读取 r 的请求主体，并将具体内容读入 body 中
	r.Body.Read(body)
	// 将字节切片内容写入相应报文
	fmt.Println("body:", string(body))
}
