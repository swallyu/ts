package ts

import "net/http"

type Context struct {
	resp   http.ResponseWriter
	Req    *http.Request
	Params Params
	app    *App
}

func (c *Context) Text(data interface{}) {

}

func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}
