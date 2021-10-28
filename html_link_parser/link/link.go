package link

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatalln("html file parse error:", err)
	}

	return getAllLinks(doc)
}

func getAllLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				text := extractText(n.FirstChild)
				links = append(links, Link{a.Val, text})
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		exLinks := getAllLinks(c)
		links = append(links, exLinks...)
	}

	return links
}

func extractText(node *html.Node) string {
	var text string
	if node.Type != html.ElementNode && node.Data != "a" && node.Type != html.CommentNode {
		text = strings.TrimSpace(node.Data)
	}

	for c := node.FirstChild; c != nil; c = node.NextSibling {
		text += extractText(c)
	}

	return strings.TrimSpace(text)

}
