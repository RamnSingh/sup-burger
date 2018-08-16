package models

import (
	"net/http"
)

type Model interface {
	PopulateFromForm(req http.Request)
}
