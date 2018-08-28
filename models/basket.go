package models

type ProductWithQuantity struct {
  Burger Burger `json:burger`
  Quantity int `json:quantity`
  Total float64 `json:total`
}

type Basket struct{
  ProductsWithQuantities [] ProductWithQuantity
  Total float64
}
