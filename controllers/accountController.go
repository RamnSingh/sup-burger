package controllers

import(
  core "../core"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "../models"
  validation "../validators"
  "strconv"
  "../utils"
  "encoding/json"
  "strings"
  mw "../middleware"
  "../helpers/form"
  // "encoding.gob"
)

type AccountController struct{}

var accountHelper *form.AccountFormHelper = &form.AccountFormHelper{}

func (controller *AccountController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/account/register", mw.NotLoggedIn(AccountRegisterHandler)).Methods("GET", "POST")
  router.HandleFunc("/account/login", mw.NotLoggedIn(AccountLoginHandler)).Methods("GET", "POST")
  router.HandleFunc("/account/logout", mw.LoggedIn(AccountLogoutHandler)).Methods("POST")
}

func AccountRegisterHandler(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
    req.ParseForm()
    user := accountHelper.PopulateFromRegisterForm(req.Form)
    if len(req.Form["cf-password"]) > 0 {
      err := validation.RegisterUser(user, req.Form["cf-password"][0], req.Form["city"][0])
      if err != nil {
        fmt.Println(err.Error())
      }else{
        cities, err := models.GetAllCities()
        if err != nil {
          panic(err.Error)
        }
        for _, city := range cities {
          if strconv.Itoa(city.ID) == req.Form["city"][0] {
            user.City = city
          }
        }
        err = user.Register(1)
        if err != nil {
          fmt.Println(err.Error())
        }else{
          http.Redirect(res, req, "/account/login", http.StatusSeeOther)
        }
      }
    }else{
      fmt.Println("confirm password not found")
    }
  }else {
    data, err := models.GetAllCities()
    if err != nil {
      core.View(res, req, "account/index.html", nil)
    } else {
      core.View(res,req, "account/register.html", data)
    }

  }
}

func AccountLoginHandler(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
    req.ParseForm()
    user := accountHelper.PopulateFromLoginForm(req.Form)
    err := validation.LoginUser(user)
    if err != nil {
      fmt.Println(err.Error())
    }else{
      err := user.Login()
      if err != nil {
        fmt.Println(err.Error())
        core.View(res, req,"account/login.html", nil)
      }else{
        userJson, err := json.Marshal(user)
        if err != nil{
          core.View(res,req, "account/login.html", nil)
        }else{
          if err = utils.SaveToSession("user", string(userJson[:]) , res, req); err != nil {
            core.View(res,req, "account/login.html", nil)
          }else{
            if strings.ToLower(user.Role.Name) == "admin" {
              http.Redirect(res, req, "/dashboard", http.StatusSeeOther)
            }else{
              http.Redirect(res, req, "/burgers", http.StatusSeeOther)
            }
          }
        }

      }
    }
  }else{
    core.View(res,req, "account/login.html", nil)
  }
}

func AccountLogoutHandler(res http.ResponseWriter, req *http.Request){
  utils.DestroySession(res, req)
  http.Redirect(res, req, "/", http.StatusSeeOther)
}
