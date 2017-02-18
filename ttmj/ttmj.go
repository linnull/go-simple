package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const per = "http://www.ttmeiju.com/meiju/"

type Data struct {
	Itmes []string
}

var urls = []string{}

var wg sync.WaitGroup

func main() {

	b, err := ioutil.ReadFile("ttmj.json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	data := &Data{[]string{}}
	err = json.Unmarshal(b, data)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for _, item := range data.Itmes {
		wg.Add(1)
		go writeFile(item)
	}

	wg.Wait()

}

func gethtml(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	doc.Find(".Scontent ").Each(func(i int, s *goquery.Selection) {
		text, _ := s.Find("a").Attr("href")
		title := s.Find("a").Text()
		urls = append(urls, fmt.Sprint("http://www.ttmeiju.com"+text+" "+title))
	})
}

func writeFile(item string) {

	defer wg.Done()

	gethtml(fmt.Sprint(per + item + ".html"))
	file, err := os.OpenFile(fmt.Sprintf("%s.txt", item), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	file.Write([]byte(fmt.Sprint(per+item+".html") + "\r\n"))
	for _, str := range urls {
		file.Write([]byte(str + "\r\n"))
	}
	urls = []string{}

}
