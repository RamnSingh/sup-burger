package controllers

import(
  core "../core"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "../models"
  "../validators/account_validator"
  "errors"
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
      err := account_validator.Register(user, req.Form["cf-password"][0])

      if err != nil {
        fmt.Println(err.Error())
      }
    }else{
      fmt.Println("confirm password")
    }
  }else {
    core.View(res, "account/register.html", nil)
  }
}

func AccountLoginHandler(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Login"))
}

func AccountLogoutHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Logout"))
}
