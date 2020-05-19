package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	dbase "../dbase"
	models "../models"
)

// GetAllPosts ...
func GetAllPosts(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	//-------ENTITY---------------------------------------------------------
	posts, err := db.SelectPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pc, err := db.SelectCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users, err := db.SelectUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//---------------------DTO---------------------------
	DTOs := []models.PostDTO{}
	for i := 0; i < len(posts); i++ {
		pDTO := models.PostDTO{}
		pDTO.ID = posts[i].ID
		for _, v := range users {
			if v.ID == posts[i].AuthorID {
				a := models.AuthorDTO{}
				a.ID = v.ID
				a.Nickname = v.Nickname
				pDTO.Author = a
			}
		}
		pDTO.Title = posts[i].Title
		pDTO.Content = posts[i].Content
		ar := []string{}
		for j := 0; j < len(pc); j++ {
			if posts[i].ID == pc[j].PostID {
				ar = append(ar, pc[j].CategoryName)
			}
		}
		pDTO.Categories = ar
		pDTO.Likes, err = db.CountReactionsToPost(1, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pDTO.Dislikes, err = db.CountReactionsToPost(0, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pDTO.Comments, err = db.CountComments(posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pDTO.CreationDate = posts[i].CreationDate
		DTOs = append(DTOs, pDTO)
	}

	SendJSON(w, &DTOs)
}

// GetPostByID ...
func GetPostByID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	p, ok := r.URL.Query()["id"]
	if !ok || len(p[0]) < 1 {
		log.Println("GetPostByID: Url Param 'id' is missing")
		http.Error(w, "Internal Server Error, please try again later", http.StatusInternalServerError)
		return
	}
	postID, err := strconv.Atoi(p[0])
	if err != nil {
		log.Println("GetPostByID Atoi:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//----------ENTITY---------------------------------------------------------------
	post, err := db.SelectPostByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//-------DTO PREPARATION---------------------------------------------------------
	postDTO := models.PostDTO{}
	postDTO.ID = post.ID
	postDTO.Author = models.AuthorDTO{}
	user, err := db.SelectUserByID(post.AuthorID)
	postDTO.Author.ID = user.ID
	postDTO.Author.Nickname = user.Nickname
	postDTO.Title = post.Title
	postDTO.Content = post.Content

	pc, err := db.SelectCategoriesByPostID(post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ar := []string{}
	for j := 0; j < len(pc); j++ {
		if post.ID == pc[j].PostID {
			ar = append(ar, pc[j].CategoryName)
		}
	}
	postDTO.Categories = ar
	postDTO.Likes, err = db.CountReactionsToPost(1, post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postDTO.Dislikes, err = db.CountReactionsToPost(0, post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postDTO.Comments, err = db.CountComments(post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postDTO.CreationDate = post.CreationDate
	SendJSON(w, &postDTO)
}

// NewPost ...
func NewPost(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.PostDTO
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post := models.Post{}
	post.AuthorID = new.Author.ID
	post.Title = new.Title
	post.Content = new.Content
	//-------STARTING TRANSACTION--------
	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Cannot start transaction")
		return
	}
	ID, err := db.CreatePost(post, tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	for i := 0; i < len(new.CategoriesID); i++ {
		err = db.AssociateCategory(ID, new.CategoriesID[i], tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}
