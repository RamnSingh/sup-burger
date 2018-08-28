package core

import(
	"html/template"
	"../utils"
	"net/http"
	"strings"
	"../models"
	"fmt"
	"encoding/json"
	"path"
)

func View(res http.ResponseWriter, req *http.Request, layout string, data interface{}) {
  tpl := template.New("template")
  tpl = template.Must(

	tpl.ParseFiles("templates/shared/layout.html","templates/shared/head.html","templates/shared/header.html","templates/shared/main.html","templates/shared/footer.html",path.Join("templates", layout)))

	userJson, _ := utils.GetFromSession("user", req)

	basketJson, _ := utils.GetFromSession("basket", req)

	var user *models.User
	var basket *models.Basket

	if len(fmt.Sprint(userJson)) > 0{
		userBlob := []byte(fmt.Sprint(userJson))
		json.Unmarshal(userBlob, &user)
	}

	if len(fmt.Sprint(basketJson)) > 0 {
		basketBlob := []byte(fmt.Sprint(basketJson))
		json.Unmarshal(basketBlob, &basket)
	}

	isLoggedIn := user != nil
	isAdmin := isLoggedIn && strings.ToLower(user.Role.Name) == "admin"

  err := tpl.ExecuteTemplate(res, "layout", map[string]interface{}{
		"session" : user,
		"isLoggedIn": isLoggedIn,
		"isAdmin" : isAdmin,
		"data" : data,
		"basket" : basket,
	})

  if err != nil {
    panic(err.Error())
  }
}
