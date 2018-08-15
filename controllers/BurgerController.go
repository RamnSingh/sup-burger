package controllers

import(
  _ "../core"
  "net/http"
  "github.com/gorilla/mux"
)

type BurgerController struct{}

func (controller *BurgerController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/burgers", BurgerListHandler).Methods("GET")
  router.HandleFunc("/burgers/{id}", BurgerDetailsHandler).Methods("GET")
  router.HandleFunc("/burgers/edit", BurgerEditHandler).Methods("POST")
  router.HandleFunc("/burgers/delete", BurgerDeleteHandler).Methods("POST")
}


func BurgerListHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("List burgers"))
}

func BurgerDetailsHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Profile burger"))
}

func BurgerEditHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Edit burger"))
}

func BurgerDeleteHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("delete burger"))
}

func BurgerAddCityHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Burger add city"))
}