package models

import "time"

type Order struct {
  ID int64 `json:"id"`
  TotalPrice float64 `json:"totalPrice" validate:"required"`
  PdfPath string `json:"pdfPath" validate:"required"`
  At time.Time `json:"at"`
  User User `json:"user"`
  Burger Burger `json:"burger"`
  City City `json:"city"`
}
