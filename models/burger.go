package models

import(
  "../db"
  "strings"
  "database/sql"
)

type Burger struct {
  ID int `json:"id"`
  Name string `json:"name" validate:"required"`
  Description string `json:"description" validate:"required"`
  ImgPath string `json:"imgPath"`
  Price float64 `json:"price" validate:"required"`
  Stock int `json:"stock" validate:"required"`
  Stuffs []Stuff `json:"stuffs"`
  Cities []City `json:"cities"`
}
func (burger *Burger)Add () error {


  query := "INSERT INTO burger (name, description, img_path, price, stock) VALUES (?,?,?,?,?)"

  args := []interface{}{burger.Name, burger.Description, burger.ImgPath, burger.Price, burger.Stock}

  res, err := db.Insert(query, args...)

  if err != nil {
    return err
  }
  int64Id, err := res.LastInsertId()
  if err == nil {
    burger.ID = int(int64Id)
  }
  return nil
}
func (burger *Burger) GetOne(id int) (error){
  query := `SELECT
            burger.id as burgerId, burger.name as burgerName,burger.description as burgerDescription,
            burger.img_path as burgerImagePath, burger.price as burgerPrice, burger.stock as burgerStock,
            stuff.name as stuffName, stuff.id as stuffId,
            city.name as cityName, city.id as cityId
            FROM burger
            LEFT JOIN burger_stuff as bs
            ON bs.burger_id = burger.id
            LEFT JOIN stuff ON stuff.id = bs.stuff_id
            LEFT JOin city_burger as cb
            ON cb.burger_id = burger.id
            LEFT JOIN city ON city.id = cb.city_id
            where burger.id = ?`

  args := []interface{}{id}

  rows, err := db.Select(query, args...)

  burgerPopulated := false

  if err != nil {
    return  err
  }

  defer rows.Close()
  stuffs := make([]Stuff, 0)
  cities := make([]City, 0)
  for rows.Next() {
    var burgerId, burgerStock int
    var burgerName, burgerDescription, burgerImagePath  string
    var stuffId, cityId sql.NullInt64
    var stuffName, cityName sql.NullString
    var burgerPrice float64

    if err := rows.Scan(&burgerId, &burgerName, &burgerDescription, &burgerImagePath, &burgerPrice, &burgerStock, &stuffName, &stuffId, &cityName, &cityId); err !=  nil {
      return err
    }

    if burgerPopulated == false {
      burger.ID = burgerId
      burger.Name = burgerName
      burger.Description = burgerDescription
      burger.Price = burgerPrice
      burger.Stock = burgerStock
      burger.ImgPath = burgerImagePath
      burgerPopulated = true
    }

    if stuffId.Valid && stuffName.Valid && len(strings.TrimSpace(stuffName.String)) > 0 {
      stuff := Stuff{
        ID : int(stuffId.Int64),
        Name : stuffName.String,
      }
      stuffs = append(stuffs, stuff)
    }

    if cityId.Valid && cityName.Valid && len(strings.TrimSpace(cityName.String)) > 0 {
      city := City{
        ID : int(cityId.Int64),
        Name : cityName.String,
      }
      cities = append(cities, city)
    }

  }
  if err != nil {
    return err
  }
  burger.Stuffs = stuffs
  burger.Cities = cities
  return nil
}
func GetAllBurgers(orderBy map[string]string) ([]Burger, error){
  selectQuery := "SELECT * FROM burger"
  orderBySlice := make([]string, 0)
  if orderBy != nil {
    if orderBy["name"] == "asc" {
      orderBySlice = append(orderBySlice, "name ASC")
    } else if orderBy["name"] == "desc" {
      orderBySlice = append(orderBySlice, "name DESC")
    }

    if orderBy["price"] == "asc" {
      orderBySlice = append(orderBySlice, "price ASC")
    } else if orderBy["price"] == "desc" {
      orderBySlice = append(orderBySlice, "price DESC")
    }
  }

  orderByArr := orderBySlice

  if len(orderByArr) > 0 {
    selectQuery += " ORDER BY " + strings.Join(orderByArr, ", ")
  }
  rows, err := db.Select(selectQuery)
  burgers := make([]Burger,0)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  for rows.Next() {
    var id, stock int
    var name, description, img_path string
    var price float64

    if err := rows.Scan(&id, &name, &description, &img_path, &price, &stock,); err !=  nil {
      return nil, err
    }
    dbBurger := Burger{
      ID : id,
      Name : name,
      Description : description,
      ImgPath : img_path,
      Price : price,
      Stock : stock,
    }
    burgers = append(burgers, dbBurger)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }
  return burgers, nil
}

func (burger *Burger) UpdateBurgerStock(stock int, id int) error {
  query := "UPDATE burger SET stock = stock - ? where id = ?"
  args := []interface{}{stock, id}
  _, err := db.Update(query, args...)

  if err != nil {
    return err
  }
  return nil
}
func (burger *Burger) Update() error {
  query := "UPDATE burger SET name = ?, description = ?, price = ?, stock = ? where id = ?"
  args := []interface{}{burger.Name, burger.Description, burger.Price, burger.Stock, burger.ID}
  _, err := db.Update(query, args...)

  if err != nil {
    return err
  }
  return nil
}

func (burger *Burger) Delete(id int) (error) {
  var err error
  query := "set foreign_key_checks = 0"
  if err = db.Exec(query); err == nil {
    query = "DELETE FROM burger WHERE id = ?"
    args := []interface{}{id}
    _, err = db.Delete(query, args...)
  }
  return err
}
