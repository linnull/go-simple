package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var URL = "http://www.163.com"

var wg sync.WaitGroup

func main() {
	links, err := getLinks(URL)
	if err != nil {
		panic(err)
	}

	os.Mkdir("data", 0777)

	for _, link := range links {
		wg.Add(1)
		go downloadImage(link)
	}
	wg.Wait()
}

func getLinks(url string) ([]string, error) {
	links := []string{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	sel := doc.Find("img")

	for _, n := range sel.Nodes {

		for _, i := range n.Attr {
			if i.Key == "src" {
				links = append(links, i.Val)
			}
		}
	}
	return links, err
}

func downloadImage(url string) error {
	defer wg.Done()
	name := getName(url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(fmt.Sprintf("data/%s", name))
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, resp.Body)
	return nil
}

func getName(url string) string {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}
