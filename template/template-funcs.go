package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

func Welcome() string {
	return "Welcome"
}

func Doing(name string) string {
	return name + ", Learning Go Web template"
}

func sayHello1(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./funcs.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}

	loveGo := func() string {
		return "欢迎学习GoWeb"
	}
	tmpl1, err := template.New("funcs").Funcs(template.FuncMap{"loveGo": loveGo}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create failed: ", err)
		return
	}

	funcMap := template.FuncMap{
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "hujizhe"
	tmpl2, err := template.New("test").Funcs(funcMap).Parse("{{Welcome}}\n{{Doing .}}\n")
	if err != nil {
		panic(err)
	}
	tmpl1.Execute(w, nil)
	tmpl2.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHello1)
	http.ListenAndServe("localhost:8080", nil)
}
