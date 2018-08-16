package models

import (
  "net/http"
  "strconv"
  "strings"
  "../db"
  "../utils"
  // "database/sql"
)

type User struct {
  ID int64 `json:"id"`
  Username string `json:"username" validate:"required"`
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
  values = req.Form["username"]
  if len(values) > 0 {
    user.Username = values[0]
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

func (user *User) Register(roleId uint) (error){
  password, err := utils.Hash(user.Password)
  if err != nil{
    return err
  }
  query := "INSERT INTO user (username, email, password, blocked, street, city, role) VALUES(?,?,?,?,?,?,?)"
  args := []interface{}{user.Username, user.Email, password, false, user.Street, user.City}

  _, err = db.Insert(query, args...)

  if err != nil {
    return err
  }
  user.Password = password
  return nil
}