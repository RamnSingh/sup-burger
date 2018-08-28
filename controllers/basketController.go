package controllers

import (
	"net/http"
	"../core"
	"github.com/gorilla/mux"
	"../helpers/form"
	"fmt"
	"../models"
	"../utils"
	"encoding/json"
	mw "../middleware"
	"github.com/jung-kurt/gofpdf"
	"time"
	"path"
	"strconv"
)

type BasketController struct{}

var basketFormHelper *form.BasketFormHelper = &form.BasketFormHelper{}

func (controller *BasketController) RegisterHandles(router *mux.Router) {
	router.HandleFunc("/basket", BasketShowHandler).Methods("GET")
	router.HandleFunc("/basket/add", BasketAddHandler).Methods("POST")
	router.HandleFunc("/basket/update", BasketUpdateHandler).Methods("POST")
	router.HandleFunc("/basket/delete", BasketDeleteHandler).Methods("POST")
	router.HandleFunc("/basket/checkout", mw.LoggedIn(BasketCheckoutHandler)).Methods("POST")
	router.HandleFunc("/basket/receipt/download", mw.LoggedIn(BasketReceiptDownloadHandler)).Methods("POST")
}

func BasketShowHandler(res http.ResponseWriter, req *http.Request) {
	core.View(res, req, "basket/index.html", basketFormHelper.GetBasket(req))
}

func BasketAddHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	productWithQuantity := basketFormHelper.PopulateFromAddToCartForm(req.Form)
	err := productWithQuantity.Burger.GetOne(productWithQuantity.Burger.ID)
	if err == nil {
		basket := basketFormHelper.AddToBasket(req, productWithQuantity)
		basketBlob, err := json.Marshal(basket)
		if err == nil {
			utils.SaveToSession("basket", string(basketBlob[:]), res, req)
		}
	}
	http.Redirect(res, req, req.Header["Referer"][0], http.StatusSeeOther)
}

func BasketUpdateHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	burgerId, quantity := basketFormHelper.PopulateFromUpdateCartForm(req.Form)

	if burgerId > 0 && quantity >= 0 {
		basket := basketFormHelper.GetBasket(req)
		for index, pq := range basket.ProductsWithQuantities {
			if pq.Burger.ID == burgerId {
				if pq.Burger.Stock >= quantity {
					quantity = quantity - basket.ProductsWithQuantities[index].Quantity
					basket.ProductsWithQuantities[index].Total += float64(quantity) * basket.ProductsWithQuantities[index].Burger.Price
					basket.ProductsWithQuantities[index].Quantity += quantity
					basket.Total += float64(quantity) * basket.ProductsWithQuantities[index].Burger.Price
					if basket.ProductsWithQuantities[index].Quantity == 0 || quantity == 0 {
						basket = basketFormHelper.DeleteProductFromBasket(req, burgerId)
					}
				}
				SaveBasketState(res, req, basket)
				break
			}

		}
	}
	http.Redirect(res, req, "/basket", http.StatusSeeOther)
}

func BasketDeleteHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	burgerId  := basketFormHelper.PopulateFromDeleteCartForm(req.Form)
	basket := basketFormHelper.DeleteProductFromBasket(req, burgerId)
	SaveBasketState(res, req, basket)
	http.Redirect(res, req, "/basket", http.StatusSeeOther)
}

func SaveBasketState(res http.ResponseWriter,req *http.Request, basket models.Basket) {
	basketBlob, err := json.Marshal(basket)
	if err == nil {
		utils.SaveToSession("basket", string(basketBlob[:]), res, req)
	}
}

func BasketCheckoutHandler(res http.ResponseWriter, req *http.Request) {
	core.View(res, req, "/basket/receipt.html", nil)
}

func BasketReceiptDownloadHandler(res http.ResponseWriter, req *http.Request){
	userJson, err := utils.GetFromSession("user", req)

	basketJson, err := utils.GetFromSession("basket", req)

	var user *models.User
	var basket *models.Basket

	if len(fmt.Sprint(userJson)) > 0{
		userBlob := []byte(fmt.Sprint(userJson))
		json.Unmarshal(userBlob, &user)
	}

	if len(fmt.Sprint(basketJson)) > 0 {
		basketBlob := []byte(fmt.Sprint(basketJson))
		json.Unmarshal(basketBlob, &basket)
	}

	if basket != nil {
		fileName := strconv.Itoa(int(time.Now().Unix()))
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		pdf.SetFont("Arial", "B", 22)
		pdf.SetX(80)
		pdf.Cell(40,7, "Good Burger")
		pdf.Ln(-1)
		pdf.SetFont("Arial", "B", 16)
		pdf.SetX(60)
		pdf.Cell(40,7, "110, Rue de Punjab, 75010 Paris")
		pdf.Ln(-1)
		pdf.SetX(0)
		pdf.Ln(-1)
		pdf.Ln(-1)
		pdf.Cell(40,7, user.Username)
		pdf.Ln(-1)
		pdf.Cell(40,7, user.Street)
		pdf.Ln(-1)
		pdf.Cell(40,7, user.City.Name)
		pdf.Ln(-1)
		pdf.Ln(-1)
	  header := []string{"Name", "Price per unit", "Quantity", "Total Price"}

		for _, h := range header {
	    pdf.Cell(40,7, h)
	  }
	  pdf.Ln(-1)

	  for index, pq := range basket.ProductsWithQuantities {
			pdf.Cell(40,6,pq.Burger.Name)
			pdf.Cell(40,6,fmt.Sprint(pq.Burger.Price))
			pdf.Cell(40,6,fmt.Sprint(pq.Quantity))
			pdf.Cell(40,6,fmt.Sprint(float64(pq.Quantity) * pq.Burger.Price))
			pdf.Ln(-1)
			basket.ProductsWithQuantities[index].Burger.UpdateBurgerStock(pq.Quantity, pq.Burger.ID)
		}
		pdf.Ln(-1)
		pdf.Cell(40,6,fmt.Sprintf("Grand total : %v", basket.Total))
		err := pdf.OutputFileAndClose(path.Join("public/assets/pdf/",fileName + ".pdf"))
		if err == nil {
			order := models.Order {
				TotalPrice : basket.Total,
				PdfPath : fileName + ".pdf",
				At : time.Now(),
				UserId : user.ID,
				User : *user,
			}
			err = models.SaveOrder(order)

			if err == nil {
				utils.RemoveFromSession("basket", res, req)
				http.ServeFile(res, req, "public/assets/pdf/" + fileName + ".pdf")
			}
		}
		fmt.Println(err)
	}

	fmt.Println(err)
}
