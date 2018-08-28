package models

import (
  "../db"
  "../utils"
  "database/sql"
  "fmt"
  "errors"
  vm "../viewsmodels"
)

type User struct {
  ID int `json:"id"`
  Username string `json:"username" validate:"required"`
  Email string `json:"email" validate:"required"`
  Password string `json:"password" validate:"required"`
  ImgPath sql.NullString `json:"imgPath"`
  Blocked bool `json:"blocked"`
  Street string `json:"street" validate:"required"`
  City City `json:"city" validate:"required"`
  Role Role `json:"role"`
}

func GetAllUsers() ([]User, error){
  query := `SELECT user.id as id, username, email,  blocked, role.id as roleId, role.name as roleName FROM user
            JOIN role ON role.id = user.role_id`

  rows, err := db.Select(query)

  users := make([]User,0)
  if err != nil {
    return users, err
  }

  defer rows.Close()

  for rows.Next() {
    var id, roleId int
    var username, email, roleName string
    var blocked bool

    if err := rows.Scan(&id, &username, &email, &blocked, &roleId, &roleName); err !=  nil {
      return users, err
    }
    dbUser := User{
      ID : id,
      Username : username,
      Email : email,
      Blocked : blocked,
      Role : Role {
        ID : roleId,
        Name : roleName,
      },
    }
    users = append(users, dbUser)
  }

  if err := rows.Err(); err != nil {
    return users, err
  }
  return users, nil
}

func (user *User) Register(roleId uint) (error){
  password, err := utils.Hash(user.Password)
  if err != nil{
    return err
  }
  query := "INSERT INTO user (username, email, password, blocked, street, city_id, role_id) VALUES(?,?,?,?,?,?,?)"
  args := []interface{}{user.Username, user.Email, password, false, user.Street, user.City.ID, roleId}

  _, err = db.Insert(query, args...)

  if err != nil {
    return err
  }
  user.Password = password
  return nil
}


func (user *User) Login() (error){
  query := "SELECT user.id as id, email, password, street, blocked, city.id as cityId, city.name as cityName, role.id as roleId, role.name as roleName FROM user JOIN city ON city.id = user.city_id JOIN role ON role.id = user.role_id WHERE username = ?"
  args := []interface{}{user.Username}

  row := db.SelectRow(query, args...)

  fmt.Println(row)
  var id, cityId, roleId int
  var email, street, password, cityName, roleName string
  // var imgPath sql.NullString
  var blocked bool

  err := row.Scan(&id, &email, &password, &street, &blocked, &cityId, &cityName, &roleId, &roleName)
  if err != nil {
    return err
  }
  if utils.CheckHash(password, user.Password){
    user.ID = id
    user.Email = email
    user.Street = street
    user.Blocked = blocked
    // user.ImgPath = imgPath

    user.City = City{cityId, cityName}
    user.Role = Role{ID : roleId, Name : roleName}

    return nil
  }else{
    return errors.New("Wrong credentials")
  }
}

func (user *User) Block() error {
  query := "UPDATE user SET blocked = CASE WHEN blocked = 0 THEN 1 ELSE 0 END where id = ?"
  args := []interface{}{user.ID}

  _, err := db.Update(query, args...)

  return err
}

func (user *User) MakeAdmin() error {
  query := "UPDATE user SET role_id = ? where id = ?"
  args := []interface{}{user.Role.ID, user.ID}

  _, err := db.Update(query, args...)

  return err
}


func GetUsersPerCity() []vm.UserCityData {
  rows, err := db.Select("SELECT sum(user.id) as users, city.Name as cityName FROM `user` JOIN city ON city.id = user.city_id group by `cityName`")
  usersPerCities := make([]vm.UserCityData,0)
  if err != nil {
    fmt.Println(err.Error())
    return usersPerCities
  }
  defer rows.Close()
  for rows.Next() {
    var users int
    var city string
    if err := rows.Scan(&users, &city); err ==  nil {
      usersPerCities = append(usersPerCities, vm.UserCityData{users, city})
    }
  }

  return usersPerCities
}

func GetNumberOfAdminsAndClients() (int, int) {
  rows, err:= db.Select("select count(user.id) as number, role.name as role from `user` JOIN role ON role.id = user.role_id group by role")
  var numberOfAdmins, numberOfClients int
  if err != nil {
    return numberOfAdmins, numberOfClients
  }
  defer rows.Close()
  for rows.Next() {
    var number int
    var role string
    err := rows.Scan(&number, &role)

    if err == nil {
      if role == "admin" {
        numberOfAdmins = number
      }else if role == "client" {
        numberOfClients = number
      }
    }
  }
  return numberOfAdmins, numberOfClients
}
