package controllers

import (
	"net/http"
	_ "../core" //
	"github.com/gorilla/mux"
)

type CityController struct{}

func (controller *CityController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/cities", CityListHandler).Methods("GET")
	router.HandleFunc("/cities/{id}", CityDetailsHandler).Methods("GET")
	router.HandleFunc("/cities/edit", CityEditHandler).Methods("POST")
	router.HandleFunc("/cities/delete", CityDeleteHandler).Methods("POST")
}

func CityListHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("List city"))
}

func CityDetailsHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Profile city"))
}

func CityEditHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Edit city"))
}

func CityDeleteHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Admin city"))
}