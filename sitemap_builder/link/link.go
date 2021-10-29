package link

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

// 因为站点地图的链接不能重复，因此使用map类型保存链接
type Links map[string]int

func Parse(r io.Reader) Links {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatalln("html file parse error:", err)
	}

	return getAllLinksFromNode(doc)
}

// 根据提供的html Node节点，获取Node节点中的所有a标签的href属性值，结果保存在map中
func getAllLinksFromNode(n *html.Node) Links {
	var links = make(map[string]int)

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if _, ok := links[a.Val]; !ok {
					links[a.Val] = 0
				}
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extLinks := getAllLinksFromNode(c)
		for k := range extLinks {
			if _, ok := links[k]; !ok {
				links[k]++
			}
		}

	}

	return links
}
