package ts

const (
	GET         = "GET"
	POST        = "POST"
	PUT         = "PUT"
	DELETE      = "DELETE"
	CONNECTIBNG = "CONNECTIBNG"
	HEAD        = "HEAD"
	OPTIONS     = "OPTIONS"
	PATCH       = "PATCH"
	TRACE       = "TRACE"
)

type RNode struct {
}

type Router struct {
	node map[string]*RNode
}

func (r *Router) Get(path string, h HandleFunc) {

}
