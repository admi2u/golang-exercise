// URLShorter练习项目的目的是，实现一个http handler,可以处理http请求，
// 当请求的path和指定的path匹配时，就跳转到特定的url；
// path映射规则可以由map变量类型提供，也可以由yaml格式的字符串提供

package main

import (
	"fmt"
	"net/http"
	"urlshort/handler"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	return mux
}

func main() {
	mux := defaultMux()

	m := make(map[string]string, 1)
	m["/baidu"] = "http://www.baidu.com"
	mapHandler := handler.MapHandler(m, mux)

	yaml := `
- path: /sina
  url: http://www.sina.com
- path: /qq
  url: http://www.qq.com
`
	yamlHandler := handler.YamlHandler(yaml, mapHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}
