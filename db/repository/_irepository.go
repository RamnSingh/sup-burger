package repository

import(
  "database/sql"
  "../../models"
)

type irepository interface {
  Get(string) *sql.rows
  Delete(string) int
  Edit(string) int
  Create(models.model) int
}
