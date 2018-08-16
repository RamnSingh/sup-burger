package models

import (
	"net/http"
)

type model interface {
	PopulateFromForm(req http.Request)
}
