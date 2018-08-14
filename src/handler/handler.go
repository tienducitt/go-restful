package handler

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	WriteHeader(w http.ResponseWriter)
	WriteBody(w http.ResponseWriter)
}

type Handler func(r *http.Request) Response

func MakeHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)
		resp.WriteBody(w)
		resp.WriteHeader(w)
	}
}

type JsonReponse struct {
	Code int
	Body Body
}

type Body struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error,omitempty"`
}

func (r *JsonReponse) WriteHeader(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	w.Header().Set("Content-Type", "application/json")
}

func (r *JsonReponse) WriteBody(w http.ResponseWriter) {
	bytes, _ := json.Marshal(r.Body)
	w.Write(bytes)
}

func Success(body interface{}) *JsonReponse {
	return &JsonReponse{Code: http.StatusOK, Body: Body{Data: body}}
}

func Error(code int, err error) *JsonReponse {
	return &JsonReponse{
		Code: code,
		Body: Body{
			Error: err,
		},
	}
}
