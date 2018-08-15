package repository

import(
  "database/sql"
  "../../models"
)

type IRepository interface {
  Get(string) *sql.rows
  Delete(string) int
  Edit(string) int
  Create(models.model) int
}
