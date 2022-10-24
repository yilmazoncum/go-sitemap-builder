package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, error := html.Parse(r)
	if error != nil {
		return nil, error
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding + msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+" ")
	}
}
