package render

import (
	"fmt"
	"io"
	"net/http"
)

var plainContentType = []string{"text/plain; charset=utf-8"}

type Text struct {
	Format string
	Data   []interface{}
}

// Render (String) writes data with custom ContentType.
func (r Text) Render(w http.ResponseWriter) error {
	return WriteString(w, r.Format, r.Data)
}

// WriteContentType (String) writes Plain ContentType.
func (r Text) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, plainContentType)
}

func WriteString(w http.ResponseWriter, format string, data []interface{}) (err error) {
	writeContentType(w, plainContentType)
	if len(data) > 0 {
		_, err = fmt.Fprintf(w, format, data...)
		return
	}
	_, err = io.WriteString(w, format)
	return
}
