package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/greennit/api"
	"github.com/greennit/error"
	"github.com/greennit/util"
	"github.com/greennit/database"
)

func main() {
	router := mux.NewRouter()
	// Add service routes
	api.AddUserRoutes(router.PathPrefix("/api/users").Subrouter())

	database.InitDB()
	// Index route
	router.Handle("/", util.AppHandler(indexHandler)).Methods("GET")
	// Configure server
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	log.Println("HTTP Server - localhost:8080 has already served  ")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	return map[string]string{"page": "index"}, nil
}
