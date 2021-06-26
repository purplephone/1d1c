package scrapper

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Baek struct {
	Description string
	Input       string
	Output      string
}

var baseURL string = "https://www.acmicpc.net/problem/"

//baekjun problem scrape
func Scrape(num string) Baek {
	//set URL
	URL := baseURL + num

	//Request the HTML page
	res, err := http.Get(URL)
	CheckErr(err)
	checkCode(res)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	//Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckErr(err)

	var b Baek
	//Find the items
	doc.Find("#problem-body").Each(func(i int, s *goquery.Selection) {
		description := strings.TrimSpace(s.Find("#problem_description").Text())
		input := strings.TrimSpace(s.Find("#problem_input").Text())
		output := strings.TrimSpace(s.Find("#problem_output").Text())
		b = Baek{description, input, output}
		//writeBaek(b)
	})
	//ba := "문제\n" + b.description + "\n\n입력\n" + b.input + "\n\n출력\n" + b.output
	return b
}

func writeBaek(b Baek) {
	file, err := os.Create("baekjun.hwp")
	CheckErr(err)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	ba := "문제\n" + b.Description + "\n\n입력\n" + b.Input + "\n\n출력\n" + b.Output
	_, wErr := w.WriteString(ba)
	CheckErr(wErr)
}

func CheckErr(err error) {
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
