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
	err := os.Remove("scrape.html")
	checkErr(err)
	term = strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	err1 := process()
	checkErr(err1)
	return c.File("scrape.html")
}

func process() error {
	file, err := os.Create("scrape.html")
	defer file.Close()
	checkErr(err)
	//wr := io.Writer(file)
	str := scrapper.Scrape(term)
	tmp1, err1 := template.New("tmp1").ParseFiles("templates/tmp1")
	checkErr(err1)
	return tmp1.ExecuteTemplate(file, "tmp1", str)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
