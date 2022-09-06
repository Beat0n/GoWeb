package main

import (
	"fmt"
	"net/http"
)

func getBody(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	http.HandleFunc("/getBody", getBody)
	http.ListenAndServe(":8080", nil)
}
