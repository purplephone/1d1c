package main

import (
	"log"
	"net/http"

	handler "github.com/purplephone/learngo/Handler"
)

var myLogger *log.Logger

func main() {
	/*fpLog, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err)
	defer fpLog.Close()
	myLogger = log.New(fpLog, "", log.LstdFlags)
	myLogger.Print("Test log")
	*/

	http.HandleFunc("/", handler.HandleHome)
	http.HandleFunc("/baek", handler.HandleBaek)
	http.HandleFunc("/scrape", handler.HandleScrape)
	http.HandleFunc("/login", handler.HandleLogin)
	http.HandleFunc("/user_login", handler.HandleUser_login)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js"))))
	http.ListenAndServe(":1323", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
