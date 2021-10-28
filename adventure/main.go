// choose your own adventure 这个项目的场景介绍：
// 给定一个json文件，存储了story的内容以及options操作；
// 读取json文件中的内容，将story的内容显示在网页上；
// 在每个story的末尾显示options操作，用户可以选择options操作跳转到其他的story；

package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	jsonFile string
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseJSON(f io.Reader) Story {
	decoder := json.NewDecoder(f)
	story := make(Story)
	if err := decoder.Decode(&story); err != nil {
		log.Fatalln("json file decode error:", err)
	}

	return story
}

// 返回一个http.HandlerFunc类型，因为HandleFunc类型实现了Handler接口，因此是一个handler处理器
func storyHandler(story Story, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSpace(r.URL.Path)
		// 默认页面的内容是intro章节内容
		if r.URL.Path == "" || r.URL.Path == "/" {
			path = "/intro"
		}

		// 根据路径获取对应的章节内容
		// 然后调用html模板，并将获得的chapter内容传入到模板中解析
		// 将解析后的html字符串写入ResponseWriter返回
		path = path[1:]
		if chapter, ok := story[path]; ok {
			err := tmpl.Execute(w, chapter)
			if err != nil {
				log.Printf("%v", err)
				http.Error(w, "Something went Wrong...", http.StatusInternalServerError)
			}
		}
	}
}

func init() {
	// 解析json文件名称
	flag.StringVar(&jsonFile, "f", "", "包含story内容的json文件名称")
	flag.Parse()
}

func main() {
	// 解析json文件内容
	f, err := os.Open(jsonFile)
	if err != nil {
		log.Fatalln("story json file open error:", err)
	}
	defer f.Close()
	story := ParseJSON(f)

	// 解析html模板
	tmpl, err := template.ParseFiles("story.html")
	if err != nil {
		log.Fatalln("template parse error:", err)
	}

	// 生成http handler处理器
	handleFunc := storyHandler(story, tmpl)

	log.Println("starting http server on 8080...")
	http.ListenAndServe(":8080", handleFunc)
}
