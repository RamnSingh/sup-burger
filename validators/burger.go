package validators

import(
  "../models"
  "strings"
  "errors"
  "strconv"
  // "fmt"
)

func AddBurger(burger models.Burger, stuffArr []string, citiesArr []string)  error{

  burger.Name = strings.TrimSpace(burger.Name)
  burger.Description = strings.TrimSpace(burger.Description)
  burger.ImgPath = strings.TrimSpace(burger.ImgPath)


  if len(burger.Name) == 0{
    return errors.New("Name is required")
  }
  if len(burger.Description) == 0{
    return errors.New("Description is required")
  }
  if burger.Price <= 0{
    return errors.New("Price is required")
  }
  if burger.Stock <= 0{
    return errors.New("Stock is required")
  }

  if len(stuffArr) == 0 {
    return errors.New("Stuff is required")
  }

  for _, stuff := range stuffArr {
    _, err := strconv.Atoi(stuff)
    if err != nil {
      return errors.New("Stuff is not valid")
    }
  }
  if len(citiesArr) == 0 {
    return errors.New("Stuff is required")
  }
  for _, city := range citiesArr {
    _, err := strconv.Atoi(city)
    if err != nil {
      return errors.New("City is not valid")
    }
  }

  return nil
}

func UpdateBurger(burger models.Burger)  error{

  burger.Name = strings.TrimSpace(burger.Name)
  burger.Description = strings.TrimSpace(burger.Description)
  burger.ImgPath = strings.TrimSpace(burger.ImgPath)


  if len(burger.Name) == 0{
    return errors.New("Name is required")
  }
  if len(burger.Description) == 0{
    return errors.New("Description is required")
  }
  if burger.ID <= 0{
    return errors.New("Price is required")
  }
  if burger.Price <= 0{
    return errors.New("Price is required")
  }
  if burger.Stock <= 0{
    return errors.New("Stock is required")
  }
  return nil
}
