package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/greennit/database"
	"github.com/greennit/error"
	"github.com/greennit/service"
	"github.com/greennit/util"
)


// UserAPI - Handle USER APIs
type UserAPI struct {
	Service *service.UserService
}

// User - Extract info from request
type User struct {
	Nickname   string    `json:"nickname"`
	Pwd 			 string  	 `json:"secret_pwd"`
	Birth      string    `json:"birth"`
	Email      string    `json:"email"`
}
// AddUserRoutes - Add routes of USER to sub router
func AddUserRoutes(sub *mux.Router) {
	var client = database.GetConnection()
	var userRepo = database.UserRepo{Client: client}
	var userService = service.UserService{Repository: &userRepo}
	var API = UserAPI{&userService}
	sub.Handle("/login", util.JsonHandler(API.login)).Methods("POST")
	sub.Handle("/registration", util.JsonHandler(API.registerUser)).Methods("POST")
	sub.Handle("/profile-image", util.JsonHandler(API.uploadUserProfileImage)).Methods("POST")
	sub.Handle("/{id}", util.JsonHandler(API.updateUserInfo)).Methods("PUT")
	sub.Handle("/{id}/pass", util.JsonHandler(API.changeUserPassword)).Methods("PUT")
	sub.Handle("/{id}/notifications", util.JsonHandler(API.getUserNotifications)).Methods("GET")
	sub.Handle("/{id}/perks", util.JsonHandler(API.getUserPerks)).Methods("GET")
}

func (api *UserAPI) login(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, &error.AppError{Error: err, Message: err.Error(), Code: 500}
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, &error.AppError{Error: err, Message: error.MsgBadRequest, Code: 400}
	}
	basicToken, err := api.Service.Login(r.Context(), user.Email, user.Pwd)
	if err != nil {
		return nil, &error.AppError{Error: err, Message: err.Error(), Code: 401}
	}
	return basicToken, nil
}

func (api *UserAPI) registerUser(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {
	// Extract data
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, &error.AppError{Error: err, Message: err.Error(), Code: 500}
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, &error.AppError{Error: err, Message: error.MsgBadRequest, Code: 400} 
	}
	// Register
	newUser, registerError := api.Service.Register(user.Nickname, user.Pwd, user.Birth, user.Email)
	if registerError != nil {
		log.Printf("UserAPI;registerUser;Error when register User %s", user.Email)
		return nil, &error.AppError{Error: registerError, Message: "UserAPI;registerUser;Failed", Code: 500}
	}

	return newUser, nil
}

func (api *UserAPI) uploadUserProfileImage(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func (api *UserAPI) updateUserInfo(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func (api *UserAPI) changeUserPassword(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func (api *UserAPI) getUserNotifications(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}

func (api *UserAPI) getUserPerks(w http.ResponseWriter, r *http.Request) (interface{}, *error.AppError) {

	return "", nil
}
