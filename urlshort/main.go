package main

import (
	"fmt"
	"golang-exercise/urlshort/handler"
	"net/http"
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
