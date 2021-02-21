package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/Muslimah94/forum-back/dbase"
	"github.com/Muslimah94/forum-back/handlers"
	"github.com/Muslimah94/forum-back/models"
)

type hendler func(db *dbase.DataBase, w http.ResponseWriter, r *http.Request)

type app struct {
	db *dbase.DataBase
}

func main() {

	funcName := "main"
	log.Printf("[%s] beginning", funcName)
	defer log.Printf("[%s] termination", funcName)

	// Initializing SQLite database
	db, err := dbase.Create("forum_database")
	if err != nil {
		log.Fatalf("[%s] dbase.Create: %v\n", funcName, err)
	}
	defer db.DB.Close()
	a := app{
		db: db,
	}

	// Setting handlers to routes
	http.HandleFunc("/api/register", a.NotRequireAuthMiddleware(handlers.RegisterLogin))
	http.HandleFunc("/api/login", a.NotRequireAuthMiddleware(handlers.LogIn))
	http.HandleFunc("/api/posts", a.NotRequireAuthMiddleware(handlers.GetAllPosts))
	http.HandleFunc("/api/addpost", a.RequireAuthMiddleware(handlers.NewPost))
	http.HandleFunc("/api/post", a.NotRequireAuthMiddleware(handlers.GetPostByID))
	http.HandleFunc("/api/postsby", a.NotRequireAuthMiddleware(handlers.GetPostsByCategory))
	http.HandleFunc("/api/comment", a.RequireAuthMiddleware(handlers.NewComment))
	http.HandleFunc("/api/comments", a.NotRequireAuthMiddleware(handlers.GetCommentsByPostID))
	http.HandleFunc("/api/categories", a.NotRequireAuthMiddleware(handlers.GetCategories))
	http.HandleFunc("/api/reaction", a.RequireAuthMiddleware(handlers.NewReaction))
	http.HandleFunc("/api/logout", a.NotRequireAuthMiddleware(handlers.LogOut))

	// Creating a custom server with TLS
	const keyURL = "./ssl/forumWebApi.key"
	const certURL = "./ssl/forumWebApi.crt"
	port := ":8080"
	cert, err := tls.LoadX509KeyPair(certURL, keyURL)
	if err != nil {
		log.Fatal(err)
	}
	s := &http.Server{
		Addr:      port,
		Handler:   nil, // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}
	log.Printf("[%s] Web API is listening to port %s...\n", funcName, port)
	log.Fatalf("[%s] ListenAndServeTLS: %v", funcName, s.ListenAndServeTLS("", ""))
}

func (a *app) RequireAuthMiddleware(myHandleFunc hendler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("logged-in_forum")
		if err == http.ErrNoCookie {
			handlers.SendJSON(w, models.Error{
				Status:      "Unauthorized",
				Description: "Please, authorize in order to continue.",
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
				Description: "Please, authorize in order to continue.",
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
