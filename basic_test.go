package ysnet

import (
	"fmt"
	"net/http"
	"testing"
)

type App struct {
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}

func Basic_Test(t *testing.T) {

	app := &App{}
	engine := Engine{}

	engine.GET("/", Index)

	fmt.Println("CDSX")

	error := http.ListenAndServe(":8080", app)

	fmt.Println(error)
}

func Index(ctx *Context) {
	ctx.Resp.Write([]byte("AASWD"))
	fmt.Fprint(ctx.Resp, "Welcome!\n")
}
