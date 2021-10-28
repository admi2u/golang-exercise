package handler

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// map中存储了路径path和url的映射，当请求的path和map的key值匹配时，就跳转到指定的url
func MapHandler(pathToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if originalURL, ok := pathToUrl[r.URL.Path]; ok {
			http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// 将yaml字符串解析到结构体中，当结构体的Path字段的值和请求的path匹配时，就跳转到指定的Url
func YamlHandler(data string, fallback http.Handler) http.HandlerFunc {
	urlMapper := []struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}{}
	err := yaml.Unmarshal([]byte(data), &urlMapper)
	if err != nil {
		log.Fatal("yaml string unmarshal error:", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, mapper := range urlMapper {
			if mapper.Path == r.URL.Path {
				http.Redirect(w, r, mapper.Url, http.StatusMovedPermanently)
			}
		}

		fallback.ServeHTTP(w, r)
	}
}
