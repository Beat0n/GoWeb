package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./example.tmpl")
	if err != nil {
		fmt.Println("parsefile failed:", err)
		return
	}
	name := "hjz"
	t.Execute(w, name)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("parsefile failed:", err)
		return
	}
	user := &UserInfo{
		Age:    22,
		Name:   "hjz",
		Gender: "man",
	}
	t.Execute(w, user)
}
