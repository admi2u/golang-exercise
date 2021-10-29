package link

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGetAllLinks(t *testing.T) {
	htmlString := `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page </a>
</body>
</html>
`
	r := strings.NewReader(htmlString)

	var links []Link
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatalln("testing error, html string parse error:", err)
	}

	links = getAllLinks(doc)
	for _, link := range links {
		if link.Href != "/other-page" {
			t.Fatalf("testing error, expected href value: \"%v\", but get: \"%v\"", "/other-page", link.Href)
		}

		if link.Text != "A link to another page" {
			t.Fatalf("testing error, expected text value: \"%v\", but get: \"%v\"", "A link to another page", link.Text)
		}
	}

	t.Log("测试用例运行成功! - getAllLinks")
}
