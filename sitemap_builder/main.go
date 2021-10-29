package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sitemap_builder/link"
	"strings"
)

// 定义SiteMap结构体，存储xml字段信息
type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

var (
	domain        string         // 需要生成sitemap的站点域名
	requestedUrls map[string]int // 已经请求的url
)

const (
	XMLHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// 根据提供的domain，获取该域名下的所有页面的a标签的链接
func getAllLinks(url string) map[string]int {
	links := make(map[string]int)
	// 如果该url已经请求过，就不再请求
	if _, ok := requestedUrls[url]; ok {
		return links
	}

	// 如果url是以 / 开头，就在请求的时候加上域名
	if strings.HasPrefix(url, "/") {
		url = domain + url
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("url request error:", err)
	}
	requestedUrls[url]++ // 将请求过的url放入 requestedUrls map中

	links = link.Parse(resp.Body) // 解析该url对应的页面，获取该页面所有链接
	resp.Body.Close()             // 因为要递归获取所有页面的a链接，这里就不要使用defer了

	// 遍历该url对应页面的所有a链接，并递归获取这些链接对应html页面的所有a链接
	for l := range links {
		extLinks := getAllLinks(l)
		for k := range extLinks {
			if _, ok := links[k]; !ok {
				links[k]++
			}
		}
	}

	return links
}

// 根据获取到的该域名下的所有链接，生成sitemap结构体
func createSiteMap(links map[string]int, domain string) *SiteMap {
	var siteMap SiteMap
	var urls []Url

	for u := range links {
		if strings.HasPrefix(u, "/") {
			u = domain + u
		}
		url := Url{u}
		urls = append(urls, url)
	}

	siteMap.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	siteMap.Urls = urls

	return &siteMap
}

func init() {
	flag.StringVar(&domain, "domain", "", "需要生成站点地图的域名")
	flag.Parse()

	requestedUrls = make(map[string]int)
}

func main() {
	links := getAllLinks(domain)
	siteMap := createSiteMap(links, domain)

	// 将sitemap结构体转换成xml格式的字符串
	output, err := xml.MarshalIndent(siteMap, "  ", "  ")
	if err != nil {
		log.Fatalln("xml marshal error:", err)
	}

	fmt.Println(XMLHeader + string(output))
}
