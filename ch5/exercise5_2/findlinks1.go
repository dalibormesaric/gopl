// Exercise 5.2: Write a function to populate a mapping from element names - p, div, span, and so on - to the number of elements with that name in an HTML document tree.

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	mapping := make(map[string]int)
	for k, v := range visit(mapping, doc) {
		fmt.Println(k, v)
	}
}

// visit appends to mapping each element data found in n and returns the result.
func visit(mapping map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		mapping[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mapping = visit(mapping, c)
	}
	return mapping
}
