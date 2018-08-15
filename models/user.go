package models

import (
  "net/http"
  "strconv"
  "strings"
)

type User struct {
  ID int64 `json:"id"`
  Name string `json:"name" validate:"required"`
  Email string `json:"email" validate:"required"`
  Password string `json:"password" validate:"required"`
  ImgPath string `json:"imgPath"`
  Blocked bool `json:"blocked"`
  Street string `json:"street" validate:"required"`
  City string `json:"city" validate:"required"`
  Role Role `json:"role"`
}

func (user *User) PopulateFromForm (req http.Request) {
  req.ParseForm()
  values := req.Form["id"]
  if len(values) > 0 {
    user.ID, _ = strconv.ParseInt(strings.TrimSpace(values[0]), 10, 64)
  }
  values = req.Form["name"]
  if len(values) > 0 {
    user.Name = values[0]
  }

  values = req.Form["email"]
  if len(values) > 0 {
    user.Email = values[0]
  }

  values = req.Form["password"]
  if len(values) > 0 {
    user.Password = values[0]
  }
}
