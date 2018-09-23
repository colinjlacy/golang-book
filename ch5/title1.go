package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := title(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// check if content-type is text/html
	if ct := resp.Header.Get("content-type"); ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
		resp.Body.Close()
		return fmt.Errorf("%s has content-type %s", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			fmt.Println(node.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

