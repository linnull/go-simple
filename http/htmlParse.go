package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.Request == nil {
		log.Fatal(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	printRoot(root)

	// log.Println(html.EscapeString("<>&'\""))
}

func printRoot(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				log.Println(a.Val)
				break
			}
		}
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		printRoot(n)
	}

}
