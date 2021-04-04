// Package component is a collection of basic HTML components.
package component

import (
	"fmt"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Text component.
type Text string

func (t Text) Render() []*html.Node {
	return []*html.Node{htmlg.Text(string(t))}
}

// List of components.
type List []htmlg.Component

func (l List) Render() []*html.Node {
	var nodes []*html.Node
	for _, c := range l {
		nodes = append(nodes, c.Render()...)
	}
	return nodes
}

// Join components and strings into a single component.
// Valid types are string, htmlg.Component. Panics on other input.
func Join(a ...interface{}) List {
	var list List
	for _, v := range a {
		switch v := v.(type) {
		case htmlg.Component:
			list = append(list, v)
		case string:
			list = append(list, Text(v))
		default:
			panic(fmt.Errorf("Join: unsupported type: %T", v))
		}
	}
	return list
}

// Link component.
type Link struct {
	Text   string
	URL    string
	NewTab bool // Open link in new tab.
}

func (l Link) Render() []*html.Node {
	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr:       []html.Attribute{{Key: atom.Href.String(), Val: l.URL}},
		FirstChild: htmlg.Text(l.Text),
	}
	if l.NewTab {
		a.Attr = append(a.Attr, html.Attribute{Key: atom.Target.String(), Val: "_blank"})
		// TODO: Add rel="noopener", see https://dev.to/ben/the-targetblank-vulnerability-by-example.
	}
	return []*html.Node{a}
}
