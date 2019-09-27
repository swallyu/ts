package ts

import (
	"log"
	"net/http"
	"testing"
)

var fileServer = http.FileServer(http.Dir("./html"))

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	fileServer.ServeHTTP(w, r)
}

func TestRouter_Get(t *testing.T) {

	mux := NewRouter()
	mux.GET("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "index2.html"
		ServeFiles(w, r)
	})

	mux.GET("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world user"))
	})

	mux.GET("/{file:.+}", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("hello world2"))
		//id := GetParam(r, "file")

		//w.Write([]byte("match user/:id ! get id:" + id))
		ServeFiles(w, r)
	})
	log.Fatal(http.ListenAndServe(":8181", mux))
}
