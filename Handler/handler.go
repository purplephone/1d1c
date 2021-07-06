package handler

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/purplephone/learngo/scrapper"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fs := template.Must(template.ParseFiles("assets/home.html"))
	fs.Execute(w, nil)
}

func HandleBaek(w http.ResponseWriter, r *http.Request) {
	fs := template.Must(template.ParseFiles("assets/baek.html"))
	fs.Execute(w, nil)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fs := template.Must(template.ParseFiles("assets/login.html"))
	fs.Execute(w, nil)
}

func HandleUser_login(w http.ResponseWriter, r *http.Request) {
	fs := template.Must(template.ParseFiles("assets/home.html"))
	fs.Execute(w, nil)
}

func HandleScrape(w http.ResponseWriter, r *http.Request) {
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
