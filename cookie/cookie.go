package main

import (
	"fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("test_cookie")
	fmt.Printf("cookie:%#v, err:%v\n", c, err)

	cookie := &http.Cookie{
		Name:   "test_cookie",
		Value:  "kdfaldfdafdfagfs",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("hello"))
}
func main() {
	http.HandleFunc("/", testHandle)
	http.ListenAndServe(":8080", nil)
}
