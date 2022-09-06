package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index1)
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error:%s", i)
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("error happened!")
}
