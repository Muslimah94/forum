package dbase

import (
	"database/sql"
	"fmt"
	"net/http"

	models "../models"
)

func AddNewCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var cat *models.Categories
	ReceiveJSON(r, &cat)

	st, err1 := db.Prepare(`INSERT INTO Categories (Name) VALUES (?)`)
	if err1 != nil {
		fmt.Println("AddNewCategory db.Prepare", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	_, err3 := st.Exec(cat.Name)
	if err3 != nil {
		fmt.Println("AddNewCategory st.Exec", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}

func AddPostCategories(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var cat *models.PostsCategories
	ReceiveJSON(r, &cat)

	st, err1 := db.Prepare(`INSERT INTO PostsCategories (PostID,CategoryID) VALUES (?,?)`)
	if err1 != nil {
		fmt.Println("AddNewPostCategory db.Prepare", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	_, err3 := st.Exec(cat.PostID, cat.CategoryID)
	if err3 != nil {
		fmt.Println("AddNewPostCategory st.Exec", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}
