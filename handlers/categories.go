package handlers

import (
	"net/http"

	"github.com/Muslimah94/forum-back/dbase"
	"github.com/Muslimah94/forum-back/models"
)

// GetCategories ...
func GetCategories(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	a, err := db.ReturnCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dto := models.CategoriesDTO{AllCategories: a}
	SendJSON(w, &dto)
}
