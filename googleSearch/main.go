package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var keys = []string{"golang", "java", "c++", "mysql"}

func main() {

	proxy := func(*http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}

	transport := &http.Transport{Proxy: proxy}

	http.DefaultClient.Transport = transport
	// resp, err := http.Get("https://www.google.com/search?q=java")

	for i := 0; i < len(keys); i++ {
		GetURLs("https://www.google.com/search?q=" + keys[i])
		fmt.Print("\n")
	}

}

//获得一个网址的所有超链接
func GetHyperlinks(url string) ([]string, error) {
	links := []string{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	sel := doc.Find("a")

	for _, n := range sel.Nodes {

		for _, i := range n.Attr {
			if i.Key == "href" {
				links = append(links, i.Val)
			}
		}
	}
	return links, nil

}

//爬虫主程序
func GetURLs(url string) {
	urls, err := GetHyperlinks(url)
	if err != nil {
		panic(err)
	}
	for _, url := range urls {
		//判断实际链接，排除谷歌本身的链接
		if strings.HasPrefix(url, "/url?q=") {
			//排除谷歌翻译链接
			if !strings.Contains(url, "google") && strings.Contains(url, "http") {
				//去除提取链接头尾多余部分
				d := strings.Index(url, "&sa")
				fmt.Println(url[7:d])
			}
		}
	}
}
