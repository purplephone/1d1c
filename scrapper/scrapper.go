package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://www.acmicpc.net/problem/"

//baekjun problem scrape
func Scrape(num string) {
	//set URL
	URL := baseURL + num

	//Request the HTML page
	res, err := http.Get(URL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	//Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	//Find the items
	doc.Find("#problem-body").Each(func(i int, s *goquery.Selection) {
		description := s.Find("#problem_description").Text()
		input := s.Find("#problem_input").Text()
		output := s.Find("#problem_output").Text()
		description = strings.TrimSpace(description)
		input = strings.TrimSpace(input)
		output = strings.TrimSpace(output)
		fmt.Println("문제\n" + description)
		fmt.Println("입력\n" + input)
		fmt.Println("출력\n" + output)
	})
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status :", res.StatusCode)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
