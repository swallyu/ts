package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}
var jsonpContentType = []string{"application/javascript; charset=utf-8"}
var jsonAsciiContentType = []string{"application/json"}

type JSON struct {
	Data interface{}
}

// IndentedJSON contains the given interface object.
type IndentedJSON struct {
	Data interface{}
}

// AsciiJSON contains the given interface object.
type AsciiJSON struct {
	Data interface{}
}

// PureJSON contains the given interface object.
type PureJSON struct {
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

// Render (IndentedJSON) marshals the given interface object and writes it with custom ContentType.
func (r IndentedJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

// WriteContentType (IndentedJSON) writes JSON ContentType.
func (r IndentedJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// Render (AsciiJSON) marshals the given interface object and writes it with custom ContentType.
func (r AsciiJSON) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	ret, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	for _, r := range string(ret) {
		cvt := string(r)
		if r >= 128 {
			cvt = fmt.Sprintf("\\u%04x", int64(r))
		}
		buffer.WriteString(cvt)
	}

	_, err = w.Write(buffer.Bytes())
	return err
}

// WriteContentType (AsciiJSON) writes JSON ContentType.
func (r AsciiJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonAsciiContentType)
}

// Render (PureJSON) writes custom ContentType and encodes the given interface object.
func (r PureJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(r.Data)
}

// WriteContentType (PureJSON) writes custom ContentType.
func (r PureJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&obj)
	return err
}
