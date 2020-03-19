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
	if data, e := handler(w, r); e != nil {
		http.Error(w, e.Message, e.Code)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
