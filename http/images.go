package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/linnull/goutil"
)

var wg sync.WaitGroup

func main() {
	URLs, _ := goutil.GetStringsFromJson("images.json")

	for i, url := range URLs {
		dirName := fmt.Sprintf("%d", i)
		os.Mkdir(dirName, 0777)
		links, err := goutil.GetImagesLinks(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, link := range links {
			wg.Add(1)
			go downloadImage(link, dirName)
		}
	}

	wg.Wait()
}

func downloadImage(url string, dirName string) error {
	defer wg.Done()

	name := getName(url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(dirName + "/" + name)
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
