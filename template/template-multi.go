package main

import (
	"log"
	"net/http"
	"text/template"
)

type UserInfo1 struct {
	Gender string
	Name   string
	Age    int
}

func main() {
	http.HandleFunc("/", tmplSample)
	http.ListenAndServe(":8080", nil)
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		log.Println("parse failed: ", err)
		return
	}
	user := &UserInfo1{
		Gender: "man",
		Name:   "hujizhe",
		Age:    22,
	}
	tmpl.Execute(w, user)
}
