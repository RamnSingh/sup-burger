package form

import(
  "net/url"
  "../../models"
  "strconv"
  "strings"
)

type AccountFormHelper struct{}

func (accountFormHelper *AccountFormHelper) PopulateFromRegisterForm (form url.Values) (user models.User) {
  values := form["id"]
  if len(values) > 0 {
    user.ID, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  values = form["username"]
  if len(values) > 0 {
    user.Username = values[0]
  }
  values = form["email"]
  if len(values) > 0 {
    user.Email = values[0]
  }
  values = form["password"]
  if len(values) > 0 {
    user.Password = values[0]
  }
  values = form["street"]
  if len(values) > 0 {
    user.Street = values[0]
  }
  return
}


func (accountFormHelper *AccountFormHelper) PopulateFromLoginForm (form url.Values) (user models.User) {

  values := form["username"]
  if len(values) > 0 {
    user.Username = values[0]
  }
  values = form["password"]
  if len(values) > 0 {
    user.Password = values[0]
  }
  return
}
