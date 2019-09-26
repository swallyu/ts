package ts

import (
	"fmt"
	"net/http"
	"sync"
)

type Param struct {
	Key   string
	Value string
}
type Params []Param

type Context struct {
	Resp   http.ResponseWriter
	Req    *http.Request
	Params Params
	app    *App
}

type HandleFunc func(ctx *Context)

func (c *Context) Text(data interface{}) {

}

type App struct {
	Router
	pool sync.Pool
}

func NewApp() *App {
	app := &App{}

	app.pool.New = func() interface{} {
		return app.allocateContext()
	}
	return app
}

func (a *App) allocateContext() *Context {
	return &Context{app: a}
}

func (e *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c := e.pool.Get().(*Context)

	c.Req = r
	c.Resp = rw
	//e.Router.ServeHTTP(rw, r)
	e.handleHttpRequest(c)
	e.pool.Put(c)
}

func (e *App) handleHttpRequest(c *Context) {
	fmt.Println("SDWER")
}
