package models

type User struct {
  ID int64 `json:"id"`
  Name string `json:"name" validate:"required"`
  ImgPath string `json:"imgPath"`
  Blocked bool `json:"blocked"`
  Street string `json:"street" validate:"required"`
  City string `json:"city" validate:"required"`
  Role Role `json:"role"`
}
