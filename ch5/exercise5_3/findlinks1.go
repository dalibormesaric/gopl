// Exercise 5.3: Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into <script> or <style> elements, since their contents are not visible in the browser.

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

// visit prints text content of a TextNode, unless it is script or style.
func visit(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(strings.TrimSpace(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.DataAtom != atom.Script && c.DataAtom != atom.Style {
			visit(c)
		}
	}
}
