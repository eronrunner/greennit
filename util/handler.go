package util

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/greennit/error"
)


type MiddlewareChain struct {
	MiddlewareHandlers []*func (w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request)
	Context *context.Context
	AppHanlder *AppHandler
}

func (chain *MiddlewareChain) Add(f ...*func (w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request)) *MiddlewareChain {
	for i := 0; i < len(chain.MiddlewareHandlers); i++ {
		chain.MiddlewareHandlers = append(chain.MiddlewareHandlers, f[i])
	}
	return chain
}

func (chain *MiddlewareChain) Handle() *http.Handler {
	for i := 0; i < len(chain.MiddlewareHandlers); i++ {
		
	}
	return nil
}

// AppHandler - basic hanlder for any requests/response before return any
type AppHandler func(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError)

func (chain *MiddlewareChain) AppHandler()

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
