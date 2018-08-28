package controllers

import(
  "../core"
  "net/http"
  "github.com/gorilla/mux"
  "../models"
  "fmt"
 "../validators"
 "strconv"
 "../helpers/form"
 "time"
 "strings"
 "../utils"
 "encoding/json"
)

type BurgerController struct{}

var burgerFormHelper *form.BurgerFormHelper = &form.BurgerFormHelper{}
var formHelper *form.FormHelper = &form.FormHelper{}

func (controller *BurgerController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/", BurgerListHandler).Methods("GET")
  router.HandleFunc("/burgers", BurgerListHandler).Methods("GET")
  router.HandleFunc("/burgers/add", BurgerAddHandler).Methods("GET", "POST")
  router.HandleFunc("/burgers/details/{id:[0-9]+}", BurgerDetailsHandler).Methods("GET")
  router.HandleFunc("/burgers/edit/{id:[0-9]+}", BurgerEditHandler).Methods("GET")
  router.HandleFunc("/burgers/edit", BurgerEditHandler).Methods("POST")
  router.HandleFunc("/burgers/delete", BurgerDeleteHandler).Methods("POST")
  router.HandleFunc("/burger/stuff/add", BurgerAddStuffHandler).Methods("POST")
  router.HandleFunc("/burger/stuff/delete", BurgerDeleteStuffHandler).Methods("POST")
  router.HandleFunc("/city/burger/add", BurgerAddCityHandler).Methods("POST")
  router.HandleFunc("/city/burger/delete", BurgerDeleteCityHandler).Methods("POST")
}

func BurgerListHandler(res http.ResponseWriter, req *http.Request){

  nameKey, _ := req.URL.Query()["name"]
  priceKey, _ := req.URL.Query()["price"]

  orderBy := make(map[string]string, 0)

  if len(nameKey) > 0 {
    orderBy["name"] =  strings.ToLower(nameKey[0])
  }

  if len(priceKey) > 0 {
    orderBy["price"] =  strings.ToLower(priceKey[0])
  }

  cities, _ := models.GetAllCities()
  stuffs, _ :=  models.GetAllStuff()
  cityKey, _ := req.URL.Query()["city"]
  stuffKey, _ := req.URL.Query()["stuff"]


  burgers, _ := models.GetAllBurgers(orderBy)
  filteredBurgers := make([]models.Burger,0)


  userJson, _ := utils.GetFromSession("user", req)
  var user *models.User
	if len(fmt.Sprint(userJson)) > 0{
		userBlob := []byte(fmt.Sprint(userJson))
		json.Unmarshal(userBlob, &user)
	}
  var cityName string
  if user != nil {
    cityName = strings.ToLower(user.City.Name)
  }else if len(cityKey) > 0 {
    cityName = strings.ToLower(cityKey[0])
  }

  citiesBurgers, _ := models.GetCitiesBurger()
  if len(cityName) > 0  && len(citiesBurgers) > 0{
    for index, burger := range burgers {
      citySlice := make([]models.City, 0)
      for _, cb := range citiesBurgers {
        if burger.ID == cb.BurgerId {
          for _, city := range cities {
            if cb.CityId == city.ID {
              citySlice = append(citySlice, city)
            }
          }
        }
      }
      burgers[index].Cities = citySlice
      for _, c := range burgers[index].Cities {
        if strings.ToLower(c.Name) == cityName {
          filteredBurgers = append(filteredBurgers, burgers[index])
        }
      }
    }
    burgers = filteredBurgers
  }



  burgersStuff, _ := models.GetBurgerStuff()
  if len(stuffKey) > 0  && len(burgersStuff) > 0{
    for index, burger := range burgers {
      stuffSlice := make([]models.Stuff, 0)
      for _, bs := range burgersStuff {
        if burger.ID == bs.BurgerId {
          for _, s := range stuffs {
            if bs.StuffId == s.ID {
              fmt.Println(1)
              stuffSlice = append(stuffSlice, s)
            }
          }
        }
      }
      burgers[index].Stuffs = stuffSlice
      for _, s := range burgers[index].Stuffs {
        if strings.ToLower(s.Name) == strings.ToLower(stuffKey[0]) {
          filteredBurgers = append(filteredBurgers, burgers[index])
        }
      }
    }
    burgers = filteredBurgers
  }



  data := map[string]interface{}{
    "Burgers" : burgers,
    "Cities" : cities,
    "Stuffs" : stuffs,
  }
  core.View(res, req, "burgers/index.html", data)
}

func BurgerAddHandler(res http.ResponseWriter, req *http.Request){
  stuff, err := models.GetAllStuff()
  cities, err := models.GetAllCities()
  if req.Method == "POST" && err == nil{
    req.ParseMultipartForm(1 * 1024 * 1024)
    burger := burgerFormHelper.PopulateFromAddForm(req.Form)
    fmt.Println(burger)
    err = validators.AddBurger(burger, req.Form["stuff"], req.Form["cities"])
    if err == nil {
      file, header, err := req.FormFile("image")
      if err == nil {
        fileNameArr := strings.Split(header.Filename, ".")
        fmt.Println(fileNameArr)
        fmt.Println(len(fileNameArr))
        if len(fileNameArr) > 1 {
          fileExtension := fileNameArr[len(fileNameArr)-1]
          fileName := strconv.Itoa(int(time.Now().Unix())) + "." + fileExtension
          fmt.Println(fileName)
          err = formHelper.UploadFile(file, "public/assets/images/burgers/" + fileName)
          if err == nil {
            burger.ImgPath = fileName
          }
        }
      }

      err = burger.Add()
      if err == nil {
        burgerStuff := make([]models.BurgerStuff, 0)
        for _, stuffOne := range stuff {
          for _, stuffFormId := range req.Form["stuff"] {
            stuffId, err := strconv.Atoi(stuffFormId)
            if err == nil && stuffOne.ID == stuffId {
              burgerStuff = append(burgerStuff, models.BurgerStuff{
                Burger : burger,
                Stuff : stuffOne,
              })
            }
          }
        }
        var burgerStuffArr models.BurgerStuffArr = burgerStuff
        err = burgerStuffArr.AddN()
        if err == nil {
          citiesBurger := make([]models.CityBurger, 0)
          for _, city := range cities {
            for _, cityFormId := range req.Form["cities"] {
              cityId, err := strconv.Atoi(cityFormId)
              if err == nil && city.ID == cityId {
                citiesBurger = append(citiesBurger, models.CityBurger{
                  City : city,
                  Burger : burger,
                })
              }
            }
          }
          var cityBurgerArr models.CityBurgerArr = citiesBurger
          err = cityBurgerArr.AddN()
          if err == nil {
            http.Redirect(res, req, "/burgers", http.StatusSeeOther)
          }
        }
      }
    }
  }



  if req.Method != "POST" || err != nil{
    data := map[string]interface{}{
      "stuff" : stuff,
      "cities" : cities,
    }
    core.View(res, req, "burgers/add.html", data)
  }
}

func BurgerDetailsHandler(res http.ResponseWriter, req *http.Request){
  vars := mux.Vars(req)
  idString := strings.TrimSpace(vars["id"])
  var err error
  if len(idString) > 0 {
    id, err := strconv.Atoi(idString)
    if err == nil {
      var burger models.Burger
      err = burger.GetOne(id)
      if err == nil {
        basket := basketFormHelper.GetBasket(req)
        for _, pq := range basket.ProductsWithQuantities {
          if pq.Burger.ID == id {
            burger.Stock  -= pq.Quantity
          }
        }
        core.View(res, req, "burgers/details.html", burger)
      }
    }
  }

  if len(idString) == 0 || err != nil {
    http.Redirect(res, req, "/burgers", http.StatusSeeOther)
  }
}

func BurgerEditHandler(res http.ResponseWriter, req *http.Request){
  stuffs, _ := models.GetAllStuff()
  cities, _ := models.GetAllCities()
  vars := mux.Vars(req)
  idString := strings.TrimSpace(vars["id"])
  var err error

  if req.Method == "GET" {
    if len(idString) > 0 {
      id, err := strconv.Atoi(idString)
      if err == nil {
        var burger models.Burger
        err = burger.GetOne(id)
        fmt.Println(burger)
        if err == nil {
          core.View(res, req, "burgers/edit.html", map[string]interface{}{
            "Burger" : burger,
            "Stuffs" : stuffs,
            "Cities" : cities,
          })
        }
      }
    }
  } else if req.Method == "POST"{
    req.ParseForm()
    burger := burgerFormHelper.PopulateFromEditForm(req.Form)
    err = validators.UpdateBurger(burger)

    if err == nil {
      err = burger.Update()
      if err == nil {
        http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
      }
    }
  }

  if err != nil {
    fmt.Println(err.Error())
    http.Redirect(res, req, "/burgers", http.StatusSeeOther)
  }
}

func BurgerDeleteHandler(res http.ResponseWriter, req *http.Request){
  req.ParseForm()
  idArr := req.Form["burger-id"]
  if len(idArr) > 0 {
    id, err := strconv.Atoi(strings.TrimSpace(idArr[0]))
    if err == nil {
      var burger models.Burger
      err = burger.GetOne(id)
      if err == nil {
        err =  burger.Delete(id)
        if err == nil {
          http.Redirect(res, req, "/burgers", http.StatusSeeOther)
        }
      }
    }
    if err != nil {
      fmt.Println(err.Error())
    }
  }
}

func BurgerAddStuffHandler(res http.ResponseWriter, req *http.Request){
  req.ParseForm()
  burgerStuff := burgerFormHelper.PopulateFromAddStuffForm(req.Form)
  if err := burgerStuff.Add(); err == nil {
    http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
  }else{
    http.Redirect(res, req, "/burgers", http.StatusSeeOther)
  }
}

func BurgerDeleteStuffHandler(res http.ResponseWriter, req *http.Request){
  req.ParseForm()
  burgerStuff := burgerFormHelper.PopulateFromDeleteStuffForm(req.Form)
  if err := burgerStuff.Delete(); err == nil {
    http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
  } else {
    http.Redirect(res, req, "/burgers", http.StatusSeeOther)
  }
}

func BurgerAddCityHandler(res http.ResponseWriter, req *http.Request){
  req.ParseForm()
  cityBurger := burgerFormHelper.PopulateFromAddCityForm(req.Form)
  if err := cityBurger.Add(); err == nil {
    http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
  } else {
    http.Redirect(res, req, "/burgers", http.StatusSeeOther)
  }
}

func BurgerDeleteCityHandler(res http.ResponseWriter, req *http.Request){
  req.ParseForm()
  cityBurger := burgerFormHelper.PopulateFromDeleteCityForm(req.Form)
  if err := cityBurger.Delete(); err == nil {
    http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
  } else {
    http.Redirect(res, req, "/burgers", http.StatusSeeOther)
  }
}
