package models

type Burger struct {
  Model
  ID int64 `json:"id"`
  Name string `json:"name" validate:"required"`
  Description string `json:"description" validate:"required"`
  ImgPath string `json:"imgPath"`
  Price float64 `json:"price" validate:"required"`
  Stock int64 `json:"stock" validate:"required"`
}
