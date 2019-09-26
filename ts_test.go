package ts

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRouter_Get(t *testing.T) {

	app := NewApp()

	app.Get("/login", Index)
	error := http.ListenAndServe(":8080", app)

	fmt.Println(error)
}

func Index(ctx *Context) {
	ctx.Resp.Write([]byte("AASWD"))
	fmt.Fprint(ctx.Resp, "Welcome!\n")
}
