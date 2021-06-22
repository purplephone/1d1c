package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/labstack/echo"
	"github.com/purplephone/learngo/scrapper"
)

func handleHome(c echo.Context) error {
	return c.File("pages/baek.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	err1 := tmp_problem(term)
	checkErr(err1)
	return c.File("pages/scrape.html")
}

func tmp_problem(term string) error {
	err := os.Remove("pages/scrape.html")
	checkErr(err)
	file, err := os.Create("pages/scrape.html")
	checkErr(err)
	defer file.Close()
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
