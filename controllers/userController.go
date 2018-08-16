package controllers

import (
	"net/http"
	_ "../core" //
	"github.com/gorilla/mux"
)

type UserController struct{}

func (controller *UserController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/users", UserListHandler).Methods("GET")
	router.HandleFunc("/users/{name}", UserProfileHandler).Methods("GET")
	router.HandleFunc("/users/edit", UserEditHandler).Methods("POST")
	router.HandleFunc("/users/makeadmin", UserMakeAdminHandler).Methods("POST")
	router.HandleFunc("/users/block", UserBlockHandler).Methods("POST")
}

func UserListHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("List user"))
}

func UserProfileHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Profile user"))
}

func UserEditHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Edit user"))
}

func UserMakeAdminHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Admin user"))
}

func UserBlockHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Block user"))
}
