package main

import(
	"fmt"
	"time"
	"net/http"
	"os"
	"io"
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

func main () {
//TODO GUIで動かす
}

func getPage(url string) {
	doc, err := goquery.NewDocument(url)
	if err !=nil {
		panic(err)
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _:= s.Attr("href")
		r := regexp.MustCompile(`http`)

		if r.MatchString(url) == true {
			find(url)
			fmt.Println(url)
			time.Sleep(time.Second * 1)
		}

	})
}

func find(url string) {
	doc, err := goquery.NewDocument(url)
	if err !=nil {
		panic(err)
	}

	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		img, _:= s.Attr("src")
		r := regexp.MustCompile(`http`)

		if r.MatchString(img) == true {
			getImage(img)

			fmt.Println(img)
			time.Sleep(time.Second * 1)
		}
	})
}

func getImage(img string) {
	res, err := http.Get(img)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	os.Mkdir("img", 0777)

	file, _:= os.Create(fmt.Sprintf("./img/download%d.jpg", res.Body))

	defer file.Close()

	io.Copy(file, res.Body)
}