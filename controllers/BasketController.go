package controllers

import (
	"net/http"
	_ "../core" //
	"github.com/gorilla/mux"
)

type BasketController struct{}

func (controller *BasketController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/basket", BasketShowHandler).Methods("POST")
	router.HandleFunc("/basket/edit", BasketUpdateHandler).Methods("POST")
	router.HandleFunc("/basket/checkout", BasketCheckoutHandler).Methods("POST")
}

func BasketShowHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Show basket"))
}

func BasketUpdateHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("update basket"))
}

func BasketCheckoutHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("checkout basket"))
}