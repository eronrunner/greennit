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
	sub.Handle("/registration", util.AppHandler(registerUser)).Methods("POST")
	sub.Handle("/profile-image", util.AppHandler(uploadUserProfileImage)).Methods("POST")
	sub.Handle("/{id}", util.AppHandler(updateUserInfo)).Methods("PUT")
	sub.Handle("/{id}/pass", util.AppHandler(changeUserPassword)).Methods("PUT")
	sub.Handle("/{id}/notifications", util.AppHandler(getUserNotifications)).Methods("GET")
	sub.Handle("/{id}/perks", util.AppHandler(getUserPerks)).Methods("GET")
}

func login(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func registerUser(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func uploadUserProfileImage(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func updateUserInfo(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func changeUserPassword(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func getUserNotifications(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func getUserPerks(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}
