package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/greennit/api"
	"github.com/greennit/error"
	"github.com/greennit/util"
)

func main() {
	router := mux.NewRouter()
	// Add service routes
	api.AddUserRoutes(router.PathPrefix("/api/users").Subrouter())

	// Index route
	router.Handle("/", util.JsonHandler(indexHandler)).Methods("GET")
	// Configure server
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	log.Printf("%s - HTTP Server - %s has already served\n\n", "Greennit", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	return map[string]string{"page": "index"}, nil
}
