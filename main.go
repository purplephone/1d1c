package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/labstack/echo"
	"github.com/purplephone/learngo/scrapper"
)

var term string

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term = strings.ToLower(scrapper.CleanString(c.FormValue("term")))

	return process()
}

func process() error {
	str := scrapper.Scrape(term)
	const tmp1 = `문제
	{{.Description}}
	입력
	{{.Input}}
	출력
	{{.Output}}
	`
	t, err := template.New("t1").Parse(tmp1)
	if err != nil {
		panic(err)
	}
	return t.Execute(os.Stdout, str)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
