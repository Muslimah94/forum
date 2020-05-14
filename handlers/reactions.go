package handlers

import (
	"fmt"
	"net/http"

	dbase "../dbase"
	models "../models"
)

// NewReaction ...
func NewReaction(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var rDTO models.ReactionDTO
	err := ReceiveJSON(r, &rDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//-------ENTITY---------------
	rea := models.Reaction{
		AuthorID:  rDTO.AuthorID,
		Type:      rDTO.Type,
		PostID:    rDTO.PostID,
		CommentID: rDTO.CommentID,
	}
	existing, err := db.SelectReaction(rea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if existing == nil {
		fmt.Println("Create reaction")
		db.CreateReaction(rea)
	} else {
		if existing.Type != rDTO.Type {
			db.UpdateReaction(rea)
		} else {
			db.DeleteReaction(rea)
		}
	}

}
