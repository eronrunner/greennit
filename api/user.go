package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/greennit/error"
	"github.com/greennit/util"
)

// AddUserRoutes - Add routes of USER to sub router
func AddUserRoutes(sub *mux.Router) {
	sub.Handle("/login", util.AppHandler(login)).Methods("POST")
	sub.Handle("/registration", util.AppHandler(register)).Methods("POST")
	sub.Handle("/profile-image", util.AppHandler(uploadProfileImage)).Methods("POST")
}

func login(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	
	return "", nil
}

func register(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	
	return "", nil
}

func uploadProfileImage(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	
	return "", nil
}
