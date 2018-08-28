package form

import(
  "net/url"
  "../../models"
  "strconv"
  "strings"
  // "fmt"
)

type BurgerFormHelper struct{}

var burgerFormHelper *BurgerFormHelper = &BurgerFormHelper{}

func (burgerFormHelper *BurgerFormHelper) PopulateFromAddForm (form url.Values) (burger models.Burger) {
  burger = models.Burger{}
  values := form["stock"]
  if len(values) > 0 {
    burger.Stock, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  values = form["price"]
  if len(values) > 0 {
    burger.Price, _ = strconv.ParseFloat(strings.TrimSpace(values[0]), 64)
  }

  values = form["name"]
  if len(values) > 0 {
    burger.Name = values[0]
  }

  values = form["description"]
  if len(values) > 0 {
    burger.Description = values[0]
  }

  values = form["img_path"]
  if len(values) > 0 {
    burger.ImgPath = values[0]
  }
  return
}

func (burgerFormHelper *BurgerFormHelper) PopulateFromEditForm (form url.Values) (burger models.Burger) {
  burger = burgerFormHelper.PopulateFromAddForm(form)
  values := form["id"]
  if len(values) > 0 {
    burger.ID, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  return
}

func (burgerFromHelper *BurgerFormHelper) PopulateFromAddStuffForm (form url.Values) (burgerStuff models.BurgerStuff) {
  burgerStuff = models.BurgerStuff{}
  values := form["burger-id"]
  if len(values) > 0 {
    burgerStuff.BurgerId, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  values = form["stuff-id"]
  if len(values) > 0 {
    burgerStuff.StuffId, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  return burgerStuff
}

func (burgerFromHelper *BurgerFormHelper) PopulateFromDeleteStuffForm (form url.Values) (burgerStuff models.BurgerStuff) {
  return burgerFormHelper.PopulateFromAddStuffForm(form)
}


func (burgerFromHelper *BurgerFormHelper) PopulateFromAddCityForm (form url.Values) (cityBurger models.CityBurger) {
  cityBurger = models.CityBurger{}
  values := form["burger-id"]
  if len(values) > 0 {
    cityBurger.BurgerId, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  values = form["city-id"]
  if len(values) > 0 {
    cityBurger.CityId, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  return cityBurger
}

func (burgerFromHelper *BurgerFormHelper) PopulateFromDeleteCityForm (form url.Values) (cityBurger models.CityBurger) {
  return burgerFormHelper.PopulateFromAddCityForm(form)
}
