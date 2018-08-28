package models

import(
  "../db"
)

type Role struct {
  ID int `json:"id"`
  Name string `json:"name" validate:"required"`
}

func GetAllRoles() ([]Role, error){
  query := "SELECT id, name FROM role"

  rows, err := db.Select(query)

  roles := make([]Role,0)
  if err != nil {
    return roles, err
  }

  defer rows.Close()

  for rows.Next() {
    var id int
    var name string

    if err := rows.Scan(&id, &name); err !=  nil {
      return roles, err
    }
    dbRole := Role{
      ID : id,
      Name : name,
    }
    roles = append(roles, dbRole)
  }

  if err := rows.Err(); err != nil {
    return roles, err
  }
  return roles, nil
}
