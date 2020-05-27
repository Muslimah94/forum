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
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
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
		id, err := GetUserIDBySession(db, r)
		if err != nil || id == 0 {
			dto.UserReaction = -1
		} else {
			reaction, err := db.SelectReaction(models.Reaction{
				AuthorID:  id,
				CommentID: comments[i].ID,
			})
			if reaction.AuthorID == 0 || err != nil {
				dto.UserReaction = -1
			} else {
				dto.UserReaction = reaction.Type
			}
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
	id, err := GetUserIDBySession(db, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//--------ENTITY---------------------------------------
	c := models.Comment{}
	c.AuthorID = id
	c.PostID = new.PostID
	c.Content = new.Content
	err = db.InsertComment(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
