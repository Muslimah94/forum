package mplx

import (
	"net/http"

	dbase "../dbase"
	handlers "../handlers"
)

// Multiplexer ...
func Multiplexer(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	p := r.URL.Path
	m := r.Method

	//----------POST---------------------------------------
	if p == "/api/posts" {
		if m == "GET" {
			handlers.GetAllPosts(db, w, r)
		} else if m == "POST" {
			handlers.NewPost(db, w, r)
		}
	} else if p == "/api/post" {
		if m == "GET" {
			handlers.GetPostByID(db, w, r)
		}
	}
	//-----------COMMENTS----------------------------------
	if p == "/api/comments" && m == "GET" {
		handlers.GetCommentsByPostID(db, w, r)
	} else if p == "/api/comment" && m == "POST" {
		handlers.NewComment(db, w, r)
	}
	//-----------Categories----------------------------------
	if p == "/api/categories" && m == "GET" {
		handlers.GetCategories(db, w, r)
	}
	//-----------RACTIONS----------------------------------
	if p == "/api/reaction" {
		if m == "POST" {
			handlers.NewReaction(db, w, r)
		} else if m == "GET" {
			handlers.FindReaction(db, w, r)
		}
	}
	//------------USER----------------------------------
	if p == "/api/register" && m == "POST" {
		handlers.RegisterLogin(db, w, r)
	}
	//------------LOGIN---------------------------------
	if p == "/api/login" && m == "POST" {
		handlers.LogIn(db, w, r)
	}
}

func Middleware(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {

}
