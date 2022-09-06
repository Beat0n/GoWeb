package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default Get"))
	})
	router.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default Post"))
	})
	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("user name: " + params.ByName("name")))
	})
	/*router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("user name: " + params.ByName("name")))
	})*/
	http.ListenAndServe(":8080", router)
}
