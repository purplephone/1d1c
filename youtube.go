package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Check_Robots() {
	URL := "https://www.youtube.com/robots.txt"

	res, err := http.Get(URL)
	checkErr(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("pre"))
	})
}
