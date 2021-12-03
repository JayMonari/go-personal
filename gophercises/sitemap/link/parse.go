package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents an anchor tag in an HTML document.
type Link struct {
	HRef string
	Text string
}

// Parse will take in an HTML document and will return a slice of links parsed
// from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link
	for _, n := range linkNodes(doc) {
		links = append(links, buildLink(n))
	}
	dfs(doc, "")
	return links, nil
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	textB := strings.Builder{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		textB.WriteString(text(c) + " ")
	}

	return strings.Join(strings.Fields(textB.String()), " ")
}

func buildLink(n *html.Node) Link {
	var l Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			l.HRef = attr.Val
			break
		}
	}
	l.Text = text(n)
	return l
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var nn []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nn = append(nn, linkNodes(c)...)
	}

	return nn
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = fmt.Sprintf("<%s>", msg)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
