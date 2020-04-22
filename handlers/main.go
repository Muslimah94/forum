package handlers

import (
	"net/http"

	dbase "../dbase"
	models "../models"
)

func GetAllPosts(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {

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
	for i := 0; i < len(posts); i++ {
		ar := []string{}
		for j := 0; j < len(pc); j++ {
			if posts[i].ID == pc[j].PostID {
				ar = append(ar, pc[j].CategoryName)
			}
		}
		posts[i].Categories = ar
		posts[i].Likes, err = db.CountReactionsToPost(1, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts[i].Dislikes, err = db.CountReactionsToPost(0, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts[i].Comments, err = db.CountComments(posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	SendJSON(w, &posts)
}

func GetPostByID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request, postID int) {
	post, err := db.SelectPost(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pc, err := db.SelectCategories()
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
	post.Categories = ar
	post.Likes, err = db.CountReactionsToPost(1, post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Dislikes, err = db.CountReactionsToPost(0, post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Comments, err = db.CountComments(post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJSON(w, &post)
}

func GetCommentsByPostID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request, postID int) {

	comments, err := db.SelectComments(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(comments); i++ {
		comments[i].Likes, err = db.CountReactionsToComment(1, comments[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments[i].Dislikes, err = db.CountReactionsToComment(0, comments[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	SendJSON(w, &comments)
}

func NewComment(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.Comments
	ReceiveJSON(r, &new)
	db.CreateComment(new)
}
