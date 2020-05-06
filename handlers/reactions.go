package handlers

import (
	"net/http"

	dbase "../dbase"
	models "../models"
)

// NewReaction ...
func NewReaction(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.ReactionDTO
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//-------ENTITY---------------
	rea := models.Reaction{
		AuthorID:  new.AuthorID,
		Type:      new.Type,
		PostID:    new.PostID,
		CommentID: new.CommentID,
	}
	n, err := db.SelectReaction(rea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if n == 0 {
		db.CreateReaction(rea)
	} else {
		db.UpdateReaction(rea)
	}

}
