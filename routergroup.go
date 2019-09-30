package ts

import "net/http"

type RouterGroup struct {
	router *Router
	app    *App
}

func DefautlRouter() *RouterGroup {
	return &RouterGroup{router: NewRouter()}
}

func (rg *RouterGroup) NewGroupRouter(group string) *RouterGroup {
	return &RouterGroup{router: rg.router.Group(group), app: rg.app}
}

func (e *RouterGroup) Get(path string, h HandleFunc) {
	e.router.GET(path, func(w http.ResponseWriter, r *http.Request) {
		e.handleRequest(w, r, h)
	})
}
func (e *RouterGroup) Post(path string, h HandleFunc) {
	e.router.POST(path, func(w http.ResponseWriter, r *http.Request) {
		e.handleRequest(w, r, h)
	})
}
func (e *RouterGroup) Delete(path string, h HandleFunc) {
	e.router.DELETE(path, func(w http.ResponseWriter, r *http.Request) {
		e.handleRequest(w, r, h)
	})
}
func (e *RouterGroup) Put(path string, h HandleFunc) {
	e.router.PUT(path, func(w http.ResponseWriter, r *http.Request) {
		e.handleRequest(w, r, h)
	})
}
func (e *RouterGroup) Patch(path string, h HandleFunc) {
	e.router.PATCH(path, func(w http.ResponseWriter, r *http.Request) {
		e.handleRequest(w, r, h)
	})
}

func (e *RouterGroup) handleRequest(w http.ResponseWriter, r *http.Request, h HandleFunc) {
	c := e.app.pool.Get().(*Context)
	c.Req = r
	c.resp = w
	h(c)
	e.app.pool.Put(c)
}

func (e *RouterGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.ServeHTTP(w, r)
}
