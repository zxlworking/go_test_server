package http_handle

import (
	"fmt"
	"net/http"
	"strings"
)

type SimpleParseResult struct {
	ParseResult map[string]string
}

type ParseResult struct {
	ParseResult map[string][]string
}

func (simpleParseResult *SimpleParseResult) parse(r *http.Request) {
	simpleParseResult.ParseResult = make(map[string]string)
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		simpleParseResult.ParseResult[k] = v[0]
	}
}

func (parseResult *ParseResult) parse(r *http.Request) {
	parseResult.ParseResult = make(map[string][]string)
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		parseResult.ParseResult[k] = v
	}
}
