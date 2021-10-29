package main

import (
	"flag"
	"fmt"
	"html_link_parser/link"
	"log"
	"os"
)

var (
	filename string
)

func init() {
	flag.StringVar(&filename, "f", "", "包含html字符串的html文件")
	flag.Parse()
}

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("html file open error:", err)
	}
	defer file.Close()

	links := link.Parse(file)
	for _, link := range links {
		fmt.Printf("link href = \"%s\"; text = \"%s\"\n", link.Href, link.Text)
	}

}
