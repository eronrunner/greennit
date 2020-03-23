package util

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/greennit/error"
)


// Middleware - Pre-handler or req/res
type Middleware func(*context.Context, http.HandlerFunc) http.HandlerFunc

// MultipleMiddleware - Implement check in addditional conditions before actually handle req/res
func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		 return h
	}

	wrapped := h
	ctx := context.TODO()
	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		 wrapped = m[i](&ctx, wrapped)
	}

	return wrapped

}

// AppHandler - basic hanlder for any requests/response before return any
type AppHandler func(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError)


// ServeHTTP - handle http req/res
func (handler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
	var code int
	var message string
	if (*r).Method == "OPTIONS" {
		return
	} else if (*r).Header.Get("Content-Type") != "application/json" {
		code = 400
		message = "Bad request"
		http.Error(w, message, code)
	} else {
		if data, err := handler(w, r); err != nil {
			code = err.Code
			http.Error(w, err.Message, err.Code)
		} else {
			code = 200
			json.NewEncoder(w).Encode(data)
		}
	}
	log.Printf("%s - [%s] - %s %d\n\n", r.Proto, r.Method, r.URL, code)
}
