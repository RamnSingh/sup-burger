package controllers

import (
	"net/http"
	"../core"
	"github.com/gorilla/mux"
	"../models"
	"strings"
	"strconv"
	"fmt"
)

type CityController struct{}

func (controller *CityController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/cities", CityListHandler).Methods("GET")
	router.HandleFunc("/cities/add", CityAddHandler).Methods("POST")
	router.HandleFunc("/cities/edit", CityEditHandler).Methods("POST")
	router.HandleFunc("/cities/delete", CityDeleteHandler).Methods("POST")
}

func CityListHandler(res http.ResponseWriter, req *http.Request) {
	cities, _ := models.GetAllCities()
	core.View(res, req, "cities/index.html", cities)
}

func CityAddHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	values := req.Form["name"]
	if len(values) > 0 && len(strings.TrimSpace(values[0])) > 0{
		var city *models.City
		city = &models.City{
			Name : strings.TrimSpace(values[0]),
		}
		err := city.Add()

		if err == nil {
			http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
		}else{
			fmt.Println(err.Error())
		}
	}else{
		fmt.Println(values)
	}
}

func CityEditHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	values := req.Form["id"]
	if len(values) > 0 {
		cityId, err := strconv.Atoi(strings.TrimSpace(values[0]))
		if err == nil {
			values = req.Form["name"]
			if len(values) > 0 && len(strings.TrimSpace(values[0])) > 0{
				cityName := strings.TrimSpace(values[0])
				var city *models.City  = &models.City{
					ID : cityId,
					Name : cityName,
				}

				err := city.Update()

				if err == nil {
					http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
				}
			}
		}
	}
}

func CityDeleteHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	values := req.Form["id"]
	if len(values) > 0 {
		if cityId, err := strconv.Atoi(strings.TrimSpace(values[0])); err == nil {
			var city *models.City = &models.City{
				ID : cityId,
			}
			if err = city.Delete(); err == nil {
				http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
			}
		}
	}
}
