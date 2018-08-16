package repository

import(
	"database/sql"
	"../../models"
)

type Repository struct {
}

func (rep *Repogsitory) Get(query string) (rows *sql.rows) {

}

func (rep *Repository) Delete(query string) (rows int) {

}

func (rep *Repository) Edit(query string) (rows int) {

}

func (rep *Repository) Create(model models.model) (rows int) {

}
