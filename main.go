package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/purplephone/learngo/scrapper"
)

var myLogger *log.Logger

func handleHome(w http.ResponseWriter, r *http.Request) {
	fs := template.Must(template.ParseFiles("assets/home.html"))
	fs.Execute(w, nil)
}

func handleBaek(w http.ResponseWriter, r *http.Request) {
	fs := template.Must(template.ParseFiles("assets/baek.html"))
	fs.Execute(w, nil)
}

func handleScrape(w http.ResponseWriter, r *http.Request) {
	term := strings.ToLower(scrapper.CleanString(r.FormValue("term")))
	if term == "" {
		fs := template.Must(template.ParseFiles("assets/baek.html"))
		fs.Execute(w, nil)
	} else {
		tmp_problem(term)
		fs := template.Must(template.ParseFiles("assets/scrape.html"))
		fs.Execute(w, nil)
	}
}

func tmp_problem(term string) {
	err := os.Remove("assets/scrape.html")
	checkErr(err)
	file, err := os.Create("assets/scrape.html")
	checkErr(err)
	defer file.Close()
	str := scrapper.Scrape(term)
	tmp1, err1 := template.New("tmp1").ParseFiles("templates/tmp1")
	checkErr(err1)
	err2 := tmp1.ExecuteTemplate(file, "tmp1", str)
	checkErr(err2)
}

func main() {
	fpLog, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err)
	defer fpLog.Close()
	myLogger = log.New(fpLog, "", log.LstdFlags)
	myLogger.Print("Test log")

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/baek", handleBaek)
	http.HandleFunc("/scrape", handleScrape)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	http.ListenAndServe(":1324", nil)

	//Check_Robots()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
