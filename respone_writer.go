package ts

import "net/http"

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type ResponeWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (w *ResponeWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	w.statusCode = defaultStatus
	w.size = noWritten
}

func (w *ResponeWriter) WriteHeaderNow(code int) {
	if !w.Written() {
		w.size = 0
		w.ResponseWriter.WriteHeader(w.statusCode)
	}
}

func (w *ResponeWriter) Status() int {
	return w.statusCode
}

func (w *ResponeWriter) Size() int {
	return w.size
}

func (w *ResponeWriter) Written() bool {
	return w.size != noWritten
}
