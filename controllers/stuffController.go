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

type StuffController struct{}

func (controller *StuffController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/stuffs", StuffListHandler).Methods("GET")
	router.HandleFunc("/stuffs/add", StuffAddHandler).Methods("POST")
	router.HandleFunc("/stuffs/edit", StuffEditHandler).Methods("POST")
	router.HandleFunc("/stuffs/delete", StuffDeleteHandler).Methods("POST")
}

func StuffListHandler(res http.ResponseWriter, req *http.Request) {
	stuffs, _ := models.GetAllStuff()
	core.View(res, req, "stuffs/index.html", stuffs)
}

func StuffAddHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	values := req.Form["name"]
	if len(values) > 0 && len(strings.TrimSpace(values[0])) > 0{
		var stuff *models.Stuff
		stuff = &models.Stuff{
			Name : strings.TrimSpace(values[0]),
		}
		fmt.Println(stuff)
		err := stuff.Add()

		if err == nil {
			http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
		}else{
			fmt.Println(err.Error())
		}
	}else{
		fmt.Println(values)
	}
}

func StuffEditHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	values := req.Form["id"]
	if len(values) > 0 {
		stuffId, err := strconv.Atoi(strings.TrimSpace(values[0]))
		if err == nil {
			values = req.Form["name"]
			if len(values) > 0 && len(strings.TrimSpace(values[0])) > 0{
				stuffName := strings.TrimSpace(values[0])
				var stuff *models.Stuff  = &models.Stuff{
					ID : stuffId,
					Name : stuffName,
				}

				err := stuff.Update()

				if err == nil {
					http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
				}
			}
		}
	}
}

func StuffDeleteHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	values := req.Form["id"]
	if len(values) > 0 {
		if stuffId, err := strconv.Atoi(strings.TrimSpace(values[0])); err == nil {
			var stuff *models.Stuff = &models.Stuff{
				ID : stuffId,
			}
			if err = stuff.Delete(); err == nil {
				http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
			}else{
				fmt.Println(err.Error())
			}
		}
	}
}
