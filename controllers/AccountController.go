package controllers

import(
  core "../core"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "../models"
)

type AccountController struct{}

func (controller *AccountController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/account/register", AccountRegisterHandler).Methods("GET", "POST")
  router.HandleFunc("/account/login", AccountLoginHandler).Methods("GET", "POST")
  router.HandleFunc("/account/logout", AccountLogoutHandler).Methods("POST")
}


func AccountRegisterHandler(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
    req.ParseForm()
    var user models.User
    user.PopulateFromForm(*req)

    if len(req.Form["cf-password"]) > 0 {
      if user.Password == req.Form["cf-password"][0] {
        
      }
    }
  }
  core.View(res, "account/register.html", nil)
}

func AccountLoginHandler(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Login"))
}

func AccountLogoutHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Logout"))
}