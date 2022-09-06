package main

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(301)
}
func main() {
	http.HandleFunc("/redirect", Redirect)
	http.ListenAndServe(":8080", nil)
}
