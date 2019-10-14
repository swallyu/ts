package ts

import (
	"net/http"
	"strings"
	"ts/binding"
	"ts/render"
)

var (
	MULTI_FORM = "multipart/form-data"
)

type Context struct {
	resp       ResponeWriter
	Req        *http.Request
	Params     Params
	app        *App
	statusCode int
}

func (c *Context) Text(code int, format string, args ...interface{}) {
	c.Render(code, render.Text{format, args})
}

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func (c *Context) JSON(code int, obj interface{}) {
	c.Render(code, render.JSON{Data: obj})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func (c *Context) AsciiJSON(code int, obj interface{}) {
	c.Render(code, render.AsciiJSON{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func (c *Context) PureJSON(code int, obj interface{}) {
	c.Render(code, render.PureJSON{Data: obj})
}

// Status sets the HTTP response code.
func (c *Context) Status(code int) {
	c.statusCode = code
}

func (c *Context) BindJSON(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.JSON)
}

//bing
func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error {
	return b.Bind(c.Req, obj)
}

func (c *Context) Render(code int, r render.Render) {

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(c.resp)
		c.resp.WriteHeaderNow(code)
		return
	}

	if err := r.Render(c.resp); err != nil {
		panic(err)
	}
}

func (c *Context) reset() {

}

func (c *Context) Path(name string) string {
	return GetParam(c.Req, name)
}

func (c *Context) Query(name string) string {
	c.parseForm()
	return c.Req.Form.Get(name)
}

func (c *Context) Form(name string) string {
	c.parseForm()
	return c.Req.PostFormValue(name)
}

func (c *Context) parseForm() {
	if strings.Contains(c.Req.Header.Get("Content-Type"), MULTI_FORM) {
		_ = c.Req.ParseMultipartForm(32 << 20)
	}
	_ = c.Req.ParseForm()
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
