package models

type Stuff struct{
  ID int64 `json:"id"`
  Name string `json:"name" validate:"required"`
  Description string `json:"description"`
  ImgPath string `json:"imgPath"`
}
