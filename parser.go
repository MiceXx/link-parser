package parser

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

//Link is a link <a href=""/> in an HTML doc
type Link struct {
	Href string
	Text string
}

var r io.Reader

//Parse will parse an html document for all link nodes
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
		}
	}
	link.Text = text(n)
	return link
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var buffer bytes.Buffer
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buffer.WriteString(text(c))
	}
	return buffer.String()
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var out []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		out = append(out, linkNodes(c)...)
	}
	return out
}
