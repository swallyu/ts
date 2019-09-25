package ysnet

import (
	"net/http"
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
}

func (c *Context) Text(data interface{}) {

}

type Engine struct {
	Router
}

func (e *Engine) New() *Engine {
	engine := &Engine{
		Router: Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
			HandleOPTIONS:          true,
		},
	}

	return engine
}

// func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	e.Router.ServeHTTP(rw, r)
// }
