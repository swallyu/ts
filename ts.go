package ts

import (
	"net/http"
	"sync"
)

type Param struct {
	Key   string
	Value string
}
type Params []Param

type HandleFunc func(ctx *Context)

type App struct {
	router      *RouterGroup
	routerGroup map[string]*RouterGroup
	pool        sync.Pool
}

func NewApp() *App {
	app := &App{router: DefautlRouter()}
	app.router.app = app

	app.pool.New = func() interface{} {
		return app.allocateContext()
	}
	return app
}

func (a *App) allocateContext() *Context {
	return &Context{app: a}
}
func (e *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.ServeHTTP(w, r)
}

func (e *App) Group(group string) *RouterGroup {
	rg := e.routerGroup[group]
	if rg == nil {
		rg = e.router.NewGroupRouter(group)
		e.routerGroup[group] = rg
	}
	return rg
}

func (e *App) Router() *RouterGroup {
	return e.router
}
