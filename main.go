// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	dbase "./dbase"
// 	"./handlers"
// )

// type hendler func(db *dbase.DataBase, w http.ResponseWriter, r *http.Request)

// type app struct {
// 	db *dbase.DataBase
// }

// func main() {
// 	db, err := dbase.Create("forumDB")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	a := app{
// 		db: db,
// 	}
// 	http.HandleFunc("/api/register", a.NotRequireAuthMiddleware(handlers.RegisterLogin))
// 	http.HandleFunc("/api/login", a.NotRequireAuthMiddleware(handlers.LogIn))
// 	http.HandleFunc("/api/posts", a.NotRequireAuthMiddleware(handlers.GetAllPosts)) // am I right ???
// 	http.HandleFunc("/api/post", a.NotRequireAuthMiddleware(handlers.NewPost))
// 	http.HandleFunc("/api/comment", a.RequireAuthMiddleware(handlers.NewComment))
// 	http.HandleFunc("/api/comments", a.NotRequireAuthMiddleware(handlers.GetCommentsByPostID))
// 	http.HandleFunc("/api/categories", a.NotRequireAuthMiddleware(handlers.GetCategories))
// 	http.HandleFunc("/api/reaction", a.RequireAuthMiddleware(handlers.NewReaction))

// 	fmt.Println("Server is listening to port :8080...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func (a *app) RequireAuthMiddleware(myHandleFunc hendler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL)
// 		url := "/api/login"
// 		UserID, err := handlers.GetUserIDBySession(db.db, r)
// 		if err != nil {
// 			http.Redirect(w, r, url, 200)
// 		}
// 		handlers.CheckCookie(r)
// 		// Check if there's my cookie
// 		// if yes {check value and compare it with db}
// 		//if there's {check expiration date} if {date is expired Delete Session, delete cookie, redirect to login} else {update expDate, handleFunc(a.db, w, r) }

// 		// else there's no cookie redirect login

// 		s, ok := r.URL.Query()["SessionID"]
// 		if !ok || len(s[0]) < 1 {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		session, err := a.db.SelectUserBySession(s[0])
// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		now := time.Now().Unix()
// 		if now >= session.ExpDate {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		myHandleFunc(a.db, w, r)
// 	}
// }

// func (a *app) NotRequireAuthMiddleware(handleFunc hendler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL)
// 		handleFunc(a.db, w, r)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	dbase "./dbase"
	mplx "./mplx"
)

func main() {
	db, err := dbase.Create("forumDB")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		mplx.Multiplexer(db, w, r)
	})

	fmt.Println("Server is listening to port :8080...")
	http.ListenAndServe(":8080", nil)
}
