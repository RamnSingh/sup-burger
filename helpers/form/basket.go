package form

import(
  "net/url"
  "../../models"
  "strconv"
  "strings"
  "fmt"
  "../../utils"
  "encoding/json"
  "net/http"
)

type BasketFormHelper struct{}

var basketFormHelper BasketFormHelper = BasketFormHelper{}

func (basketFormHelper *BasketFormHelper) PopulateFromAddToCartForm (form url.Values) (productWithQuantity models.ProductWithQuantity) {
  productWithQuantity = models.ProductWithQuantity{}
  productWithQuantity.Burger = models.Burger{}

  values := form["burger_id"]
  if len(values) > 0 {
    productWithQuantity.Burger.ID, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }

  values = form["quantity"]
  if len(values) > 0 {
    productWithQuantity.Quantity, _ = strconv.Atoi(strings.TrimSpace(values[0]))
  }
  return
}

func (basketFormHelper *BasketFormHelper) AddToBasket (req *http.Request, productWithQuantity models.ProductWithQuantity) (models.Basket) {

  basket := basketFormHelper.GetBasket(req)

  basketSlice := basket.ProductsWithQuantities[:]

  productAlreadyAdded := false
  for index, pq := range basketSlice {
    if pq.Burger.ID == productWithQuantity.Burger.ID {
      basketSlice[index].Quantity += productWithQuantity.Quantity
      basketSlice[index].Total = float64(pq.Quantity) * pq.Burger.Price
      basket.Total += basketSlice[index].Total - (float64(productWithQuantity.Quantity) * pq.Burger.Price)
      productAlreadyAdded = true;
      break;
    }
  }

  if productAlreadyAdded == false {
    productWithQuantity.Total = productWithQuantity.Burger.Price * float64(productWithQuantity.Quantity)
    basket.Total += productWithQuantity.Total
    basketSlice = append(basketSlice, productWithQuantity)
  }
  basket.ProductsWithQuantities = basketSlice

  return basket
}

func (basketFormHelper *BasketFormHelper) GetBasket (req *http.Request) (models.Basket) {
  basketJson, _ := utils.GetFromSession("basket", req)

  var basket *models.Basket
  if len(strings.TrimSpace(fmt.Sprint(basketJson))) > 0{
    basketBlob := []byte(fmt.Sprint(basketJson))
    json.Unmarshal(basketBlob, &basket)
  }

  if basket == nil {
    basket = &models.Basket{}
  }

  basket.Total = 0
  for index, pq := range basket.ProductsWithQuantities {
    basket.ProductsWithQuantities[index].Total = float64(pq.Quantity) * pq.Burger.Price
    basket.Total += basket.ProductsWithQuantities[index].Total
  }

  return *basket
}

func (basketFormHelper *BasketFormHelper) PopulateFromUpdateCartForm(form url.Values) (int, int) {
  value := strings.TrimSpace(form["burger_id"][0])
  var burgerId, quantity int
  if len(value) > 0 {
    burgerId, _ = strconv.Atoi(value)
  }
  value = strings.TrimSpace(form["quantity"][0])
  if len(value) > 0 {
    quantity, _ = strconv.Atoi(value)
  }

  return burgerId, quantity
}

func (basketFormHelper *BasketFormHelper) DeleteProductFromBasket(req *http.Request, burgerId int) (models.Basket) {
  basket := basketFormHelper.GetBasket(req)
  productsSlice := basket.ProductsWithQuantities[:]

  for index, pq := range basket.ProductsWithQuantities {
    if pq.Burger.ID == burgerId {
      basket.Total -= pq.Total
      productsSlice = append(productsSlice[:index], productsSlice[index+1:]...)
      break
    }
  }

  basket.ProductsWithQuantities = productsSlice

  return basket
}

func (basketFormHelper *BasketFormHelper) PopulateFromDeleteCartForm(form url.Values) int {
  value := strings.TrimSpace(form["burger_id"][0])
  fmt.Println(value)
  var burgerId int
  if len(value) > 0 {
    burgerId, _ = strconv.Atoi(value)
  }
  return burgerId
}
