package models

import(
  "time"
  "../db"
  vm "../viewsmodels"
)

type Order struct {
  ID int `json:"id"`
  TotalPrice float64 `json:"totalPrice" validate:"required"`
  PdfPath string `json:"pdfPath" validate:"required"`
  At time.Time `json:"at"`
  UserId int `json:userId`
  User User `json:"user"`
}

func SaveOrder(order Order) error {
  query := "INSERT INTO `order` (total_price, pdf_path, at, user_id) VALUES (?,?,?,?)"

  args := []interface{}{order.TotalPrice, order.PdfPath, order.At, order.UserId}

  _, err := db.Insert(query, args...)

  return err
}

func GetAllOrdersByUserId(userId int) ([]Order, error){
  rows, err := db.Select("SELECT * FROM `order` WHERE user_id = ?", []interface{}{userId}...)
  orders := make([]Order,0)
  if err != nil {
    return orders, err
  }

  defer rows.Close()

  for rows.Next() {
    var id, user_id int
    var total_price float64
    var at time.Time
    var pdf_path string
    if err := rows.Scan(&id, &at, &pdf_path, &total_price, &user_id); err !=  nil {
      return nil, err
    }
    dbOrder := Order{id, total_price, pdf_path, at, user_id, User{}}
    orders = append(orders, dbOrder)
  }

  if err := rows.Err(); err != nil {
    return orders, err
  }
  return orders, nil
}

func GetOrdersPerMonth () [] vm.OrderData{
  rows, err := db.Select("SELECT count(id) as orders, extract(MONTH FROM at) as month FROM sup_burger.`order` group by `month`")
  ordersSlice := make([]vm.OrderData,0)
  if err != nil {
    return ordersSlice
  }
  defer rows.Close()
  for rows.Next() {
    var orders, month int
    if err := rows.Scan(&orders, &month); err ==  nil {
      ordersSlice = append(ordersSlice, vm.OrderData{month, orders})
    }

  }
  return ordersSlice
}

func GetMoneyPerMOnth () [] vm.MoneyData{
  rows, err := db.Select("SELECT sum(total_price) as amount, extract(MONTH FROM at) as month FROM sup_burger.`order` group by `month`")
  moneyData := make([]vm.MoneyData,0)
  if err != nil {
    return moneyData
  }
  defer rows.Close()
  for rows.Next() {
    var month int
    var amount float64
    if err := rows.Scan(&amount, &month); err ==  nil {
      moneyData = append(moneyData, vm.MoneyData{ month, amount})
    }

  }
  return moneyData
}
