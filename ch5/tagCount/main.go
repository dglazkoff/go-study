package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("https://golang.org")
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(count(map[string]int{}, doc))
}

func count(tagsCount map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		tagsCount[n.Data] += 1
	}
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(tagsCount, c)
	}

	return tagsCount
}
