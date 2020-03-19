package util

import (
	"encoding/json"
	"net/http"

	"github.com/greennit/error"
)

// AppHandler - basic hanlder for any requests/response before return any
type AppHandler func(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError)

// ServerHTTP - handle http req/res
func (handler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*r).Method == "OPTIONS" {
		return
	} else if (*r).Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Bad request", 400)
	}
	if data, e := handler(w, r); e != nil {
		http.Error(w, e.Message, e.Code)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
