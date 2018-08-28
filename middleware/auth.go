package middleware

import(
	"../utils"
	"encoding/json"
	"../models"
	"errors"
	"net/http"
	"strings"
	"fmt"
)

func OnlyAdmin(){

}

func Admin(h http.HandlerFunc) http.HandlerFunc{
	return func (res http.ResponseWriter, req *http.Request){
		user, err := getUserFromSession(req)
		if(err != nil || strings.ToLower(user.Role.Name) != "admin"){
			http.Redirect(res, req, "/", http.StatusSeeOther)
		}else{
			h(res, req)
		}
	}
}

func LoggedIn(h http.HandlerFunc) http.HandlerFunc{
	return func (res http.ResponseWriter, req *http.Request){
		user, _ := getUserFromSession(req)

		if len(strings.TrimSpace(user.Username)) > 0 {
			h(res, req)
		}else{
			http.Redirect(res, req, "/account/login", http.StatusSeeOther)
		}
	}
}

func NotLoggedIn(h http.HandlerFunc) http.HandlerFunc{
	return func (res http.ResponseWriter, req *http.Request){
		user, _ := getUserFromSession(req)
		if len(strings.TrimSpace(user.Username)) ==  0{
			h(res, req)
		}else{
			http.Redirect(res, req, "/", http.StatusSeeOther)
		}

	}
}

func getUserFromSession(req *http.Request) (models.User, error){

	userJson, _ := utils.GetFromSession("user", req)

	var user models.User

	if len(fmt.Sprint(userJson)) > 0{

		userBlob := []byte(fmt.Sprint(userJson))
		if err := json.Unmarshal(userBlob, &user); err != nil {
			return user, err
		}else {
			return user, nil
		}
	}else{
		return user, errors.New("No data found")
	}
}
