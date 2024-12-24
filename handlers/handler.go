package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type Handler interface {
	ConfigureSelf(m *http.ServeMux) *http.ServeMux
}

func InitializeHandlers(db *gorm.DB) []Handler {
	return []Handler{
		newUserHandler(db),
		newQuizHandler(db),
	}
}
