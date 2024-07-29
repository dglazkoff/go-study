package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	words, images, _ := CountWordsAndImages("https://golang.org")

	fmt.Println(words)
	fmt.Println(images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	return countWords(0, n), countImages(0, n)
}

func countWords(words int, n *html.Node) int {
	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
		fmt.Println(words)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words = countWords(words, c)
	}

	return words
}

func countImages(images int, n *html.Node) int {
	if n.Type == html.ElementNode && n.Data == "img" {
		images += 1
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		images = countImages(images, c)
	}

	return images
}
