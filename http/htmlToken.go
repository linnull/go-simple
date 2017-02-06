package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.ttmeiju.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.Request == nil {
		log.Fatal(err)
	}

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			print(t)
		}
	}
}

func print(t html.Token) {
	if t.Data == "a" {
		for _, a := range t.Attr {
			if a.Key == "href" {
				log.Println("Found href:", a.Val)
				break
			}
		}
	}

}
