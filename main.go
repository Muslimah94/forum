package main

import (
	"fmt"
	"log"
	"net/http"

	dbase "./dbase"
	"./handlers"
	models "./models"
)

type hendler func(db *dbase.DataBase, w http.ResponseWriter, r *http.Request)

type app struct {
	db *dbase.DataBase
}

func main() {
	db, err := dbase.Create("forumDB")
	if err != nil {
		log.Fatal(err)
	}
	a := app{
		db: db,
	}
	http.HandleFunc("/api/register", a.NotRequireAuthMiddleware(handlers.RegisterLogin))
	http.HandleFunc("/api/login", a.NotRequireAuthMiddleware(handlers.LogIn))
	http.HandleFunc("/api/posts", a.NotRequireAuthMiddleware(handlers.GetAllPosts))
	http.HandleFunc("/api/addpost", a.RequireAuthMiddleware(handlers.NewPost))
	http.HandleFunc("/api/post", a.NotRequireAuthMiddleware(handlers.GetPostByID))
	http.HandleFunc("/api/comment", a.RequireAuthMiddleware(handlers.NewComment))
	http.HandleFunc("/api/comments", a.NotRequireAuthMiddleware(handlers.GetCommentsByPostID))
	http.HandleFunc("/api/categories", a.NotRequireAuthMiddleware(handlers.GetCategories))
	http.HandleFunc("/api/reaction", a.RequireAuthMiddleware(handlers.NewReaction))

	fmt.Println("Server is listening to port :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *app) RequireAuthMiddleware(myHandleFunc hendler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		cookie, err := r.Cookie("logged-in_forum")
		if err == http.ErrNoCookie {

			handlers.SendJSON(w, models.Error{
				Status:      "Unauthorized",
				Description: "Please authorize",
			})
			return
		}
		session, err := a.db.SelectSession(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if session.ID == 0 {
			handlers.DeleteCookie(w, r)
			handlers.SendJSON(w, models.Error{
				Status:      "Unauthorized",
				Description: "Please authorize",
			})
			return
		}
		if a.db.CompareExpDate(session) {
			tx, err := a.db.DB.Begin()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = a.db.UpdateSessionDate(session, tx)
			if err != nil {
				tx.Rollback()
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tx.Commit()
			myHandleFunc(a.db, w, r)
		} else {
			a.db.DeleteSession(session.UUID.String())
			handlers.DeleteCookie(w, r)
			handlers.SendJSON(w, models.Error{
				Status:      "Unauthorized",
				Description: "Please authorize",
			})
			return
		}
	}
}

func (a *app) NotRequireAuthMiddleware(handleFunc hendler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		handleFunc(a.db, w, r)
	}
}
