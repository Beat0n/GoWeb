package main

import (
	"fmt"
	"net/http"
	"strings"
)

func request(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request parse")

	fmt.Println("Method", r.Method)

	fmt.Println("RequestURI:", r.RequestURI)

	fmt.Println("URL_Path:", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)

	fmt.Println("Proto", r.Proto)
	fmt.Println("protomajor", r.ProtoMajor)
	fmt.Println("protominor", r.ProtoMinor)

	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println("header key:" + k + "  value:" + vv)
		}
	}

	isMultipart := false
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			isMultipart = true
		}
	}

	if isMultipart == true {
		r.ParseMultipartForm(128)
		fmt.Println("Parse method: ParseMultipartForm")
	} else {
		fmt.Println("Parse method: ParseForm")
	}

	fmt.Println("ContentLength", r.ContentLength)

	fmt.Println("Close", r.Close)

	fmt.Println("host", r.Host)

	fmt.Println("RemoteAddr", r.RemoteAddr)

	fmt.Fprintf(w, "hello let's go")
}
func main() {
	http.HandleFunc("/hello", request)
	http.ListenAndServe(":8080", nil)
}
