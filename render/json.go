package render

import (
	"encoding/json"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}
var jsonpContentType = []string{"application/javascript; charset=utf-8"}
var jsonAsciiContentType = []string{"application/json"}

type JSON struct {
	Data interface{}
}

func (r JSON) Render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&obj)
	return err
}
