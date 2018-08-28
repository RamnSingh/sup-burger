package controllers

import (
	"net/http"
	"../core"
	"github.com/gorilla/mux"
	"../models"
	"strings"
	"fmt"
	"strconv"
	mw "../middleware"
	"encoding/json"
	"../utils"
)

type UserController struct{}

func (controller *UserController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/users", mw.Admin(UserListHandler)).Methods("GET")
	router.HandleFunc("/users/profile", mw.LoggedIn(UserProfileHandler)).Methods("GET")
	router.HandleFunc("/users/orders", mw.LoggedIn(UserOrdersHandler)).Methods("GET")
	router.HandleFunc("/users/invoice", mw.LoggedIn(UserInvoiceHandler)).Methods("POST")
	router.HandleFunc("/users/block", mw.Admin(UserBlockHandler)).Methods("POST")
	router.HandleFunc("/users/makeadmin", mw.Admin(UserMakeAdminHandler)).Methods("POST")
}

func UserListHandler(res http.ResponseWriter, req *http.Request) {
	users, _ := models.GetAllUsers()
	core.View(res, req, "/users/index.html", users)
}

func UserProfileHandler(res http.ResponseWriter, req *http.Request) {
	core.View(res, req, "/users/profile.html", nil)
}

func UserOrdersHandler(res http.ResponseWriter, req *http.Request) {
	userJson, _ := utils.GetFromSession("user", req)
	var user *models.User
	if len(fmt.Sprint(userJson)) > 0{
		userBlob := []byte(fmt.Sprint(userJson))
		json.Unmarshal(userBlob, &user)
	}
	var orders []models.Order
	if user != nil {
		orders, _ = models.GetAllOrdersByUserId(user.ID)
	}
	core.View(res, req, "/users/orders.html", orders)
}

func UserInvoiceHandler(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	if len(req.Form["pdf-path"]) > 0 && len(strings.TrimSpace(req.Form["pdf-path"][0])) > 0 {
		http.ServeFile(res, req, "public/assets/pdf/" + strings.TrimSpace(req.Form["pdf-path"][0]))
	}
}

func UserBlockHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["id"]) > 0 {
		if id, err := strconv.Atoi(strings.TrimSpace(req.Form["id"][0])); err == nil {
			user := &models.User{
				ID : id,
			}
			err = user.Block()
			if err != nil {
				fmt.Println(err.Error())
			}
			http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
		}
	}
}

func UserMakeAdminHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["id"]) > 0 && len(req.Form["role-id"]) > 0  {
		if id, err := strconv.Atoi(strings.TrimSpace(req.Form["id"][0])); err == nil {
			if roleId, err := strconv.Atoi(strings.TrimSpace(req.Form["role-id"][0])); err == nil {
				roles, err := models.GetAllRoles()
				if err == nil {
					var clientId, adminId int
					for _, role := range roles {
						if strings.ToLower(strings.TrimSpace(role.Name)) == "admin" {
							adminId = role.ID
						}
						if strings.ToLower(strings.TrimSpace(role.Name)) == "client" {
							clientId = role.ID
						}
					}
					var newRoleId int
					if roleId == clientId {
						newRoleId = adminId
					}else if roleId == adminId {
						newRoleId = clientId
					}

					user := models.User{
						ID : id,
						Role : models.Role{
							ID : newRoleId,
						},
					}
					err = user.MakeAdmin()
				}
				if err != nil {
					fmt.Println(err.Error())
				}
				http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
			}
		}
	}
}
