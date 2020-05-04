package handlers

import (
	"log"
	"net/http"
	"strconv"

	dbase "../dbase"
	models "../models"
)

// GetCommentsByPostID ...
func GetCommentsByPostID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	p, ok := r.URL.Query()["post_id"]
	if !ok || len(p[0]) < 1 {
		log.Println("GetCommentsByPostID: Url Param 'post_id' is missing")
		http.Error(w, "Internal Server Error, please try again later", http.StatusInternalServerError)
		return
	}
	postID, err := strconv.Atoi(p[0])
	if err != nil {
		log.Println("GetCommentsByPostID Atoi: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//----------ENTITY---------------------------------------------------------------
	comments, err := db.SelectComments(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//-------DTO PREPARATION---------------------------------------------------------
	cDTOs := []models.CommentDTO{}
	for i := 0; i < len(comments); i++ {
		dto := models.CommentDTO{}
		dto.ID = comments[i].ID
		// ENTITY {
		user, err := db.SelectUserByID(comments[i].AuthorID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// }
		dto.Author = models.AuthorDTO{}
		dto.Author.ID = user.ID
		dto.Author.Nickname = user.Nickname
		dto.PostID = comments[i].PostID
		dto.Content = comments[i].Content
		dto.Likes, err = db.CountReactionsToComment(1, comments[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dto.Dislikes, err = db.CountReactionsToComment(0, comments[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cDTOs = append(cDTOs, dto)
	}
	SendJSON(w, &cDTOs)
}

// NewComment ...
func NewComment(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.CommentDTO
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//--------ENTITY---------------------------------------
	c := models.Comment{}
	c.AuthorID = new.Author.ID
	c.PostID = new.PostID
	c.Content = new.Content
	db.CreateComment(c)
}
