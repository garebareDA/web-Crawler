package main

import(
	"fmt"
	"time"
	"net/http"
	"os"
	"io"
	"bufio"
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

func main () {
	urlinput()
}

func urlinput() {
	fmt.Println("画像を取得するURLを入力してください")
	urlInput := bufio.NewScanner(os.Stdin)
	urlInput.Scan()
	url := urlInput.Text()

	urlMatch := regexp.MustCompile(`^http`)

	if urlMatch.MatchString(url) == true {
		inputAnsewer(url)
	} else {
		fmt.Println("正しいURLを入力してください")
		time.Sleep(time.Second * 1)
		urlinput()
	}
}

func inputAnsewer(inputURL string) {
	fmt.Println("ページの先のURLのみから画像を取得しますか？(yes/no)")
	fmt.Println("yes(指定されたURLのページ中にあるURLの先から画像を取得します)")
	fmt.Println("no(指定されたURLのページのみから画像を取得します)")
	yesOrNo := bufio.NewScanner(os.Stdin)
	yesOrNo.Scan()
	answer := yesOrNo.Text()

	answerMatchYes := regexp.MustCompile(`yes`)
	answerMatchNo := regexp.MustCompile(`no`)

	if answerMatchYes.MatchString(answer) == true {
		fmt.Println("yes")
		getPage(inputURL)

	} else if answerMatchNo.MatchString(answer) == true {
		fmt.Println("no")
		find(inputURL)

	} else {
		inputAnsewer(inputURL)
	}
}

func getPage(url string) {
	doc, err := goquery.NewDocument(url)
	if err !=nil {
		fmt.Println("正しいURLを入力してください")
		time.Sleep(time.Second * 1)
		urlinput()
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
		fmt.Println("正しいURLを入力してください")
		time.Sleep(time.Second * 1)
		urlinput()
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