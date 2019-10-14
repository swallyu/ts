package ts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

var fileServer = http.FileServer(http.Dir("./html"))

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	fileServer.ServeHTTP(w, r)
}

func TestApp_Get(t *testing.T) {

	app := NewApp()

	app.Router().Post("/user", func(ctx *Context) {

		fmt.Println(ctx.Req.Header.Get("Content-Type"))
		value := ctx.Form("name")
		fmt.Println("value is :" + value)

		postData := make(map[string]interface{})

		_ = ctx.BindJSON(&postData)

		data := make(map[string]interface{})
		data["code"] = 200
		data["msg"] = "messages"
		data["data"] = []string{"ss", "bb"}

		ctx.JSON(200, data)
	})
	app.Router().Get("/{file:.+}", func(ctx *Context) {
		ServeFiles(ctx.resp, ctx.Req)
	})
	error := http.ListenAndServe(":8888", app)
	if error != nil {
		fmt.Println(error.Error())
	}
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

func TestHttpInvoke(t *testing.T) {

	url := "http://rdctest.rastyletech.com:2222/api/getdata/devsensordata"

	postData := make(map[string]interface{})

	postData["deviceList"] = []string{"00124B001D78D48A"}

	by, _ := json.Marshal(postData)

	fmt.Println(string(by))

	resp, err := http.Post(url, "application/json", strings.NewReader(string(by)))
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(ret))
}
