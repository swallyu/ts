package binding

import "net/http"

var (
	JSON = jsonBinding{}
)

type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}
