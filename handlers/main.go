package handlers

import (
	"log"
	"net/http"
	"strconv"

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
	users, err := db.SelectUsers()

	//---------------------DTO---------------------------
	postDTOs := []models.PostDTO{}
	for i := 0; i < len(posts); i++ {
		p := models.PostDTO{}
		p.ID = posts[i].ID
		for _, v := range users {
			if v.ID == posts[i].AuthorID {
				a := models.AuthorDTO{}
				a.ID = v.ID
				a.Nickname = v.Nickname
				p.Author = a
			}
		}
		p.Title = posts[i].Title
		p.Content = posts[i].Content
		ar := []string{}
		for j := 0; j < len(pc); j++ {
			if posts[i].ID == pc[j].PostID {
				ar = append(ar, pc[j].CategoryName)
			}
		}
		p.Categories = ar
		p.Likes, err = db.CountReactionsToPost(1, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.Dislikes, err = db.CountReactionsToPost(0, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.Comments, err = db.CountComments(posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.CreationDate = posts[i].CreationDate
		postDTOs = append(postDTOs, p)
	}

	SendJSON(w, &postDTOs)
}

func GetPostByID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	p, ok := r.URL.Query()["id"]
	if !ok || len(p[0]) < 1 {
		log.Println("GetPostByID: Url Param 'id{post}' is missing")
		http.Error(w, "No request parameters found", http.StatusInternalServerError)
		return
	}
	postID, err := strconv.Atoi(p[0])
	if err != nil {
		log.Println("GetPostByID: Cannot convert string to int")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := db.SelectPost(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	SendJSON(w, &postDTO)
}

func GetCommentsByPostID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request, postID int) {

	// 	comments, err := db.SelectComments(postID)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	for i := 0; i < len(comments); i++ {
	// 		comments[i].Likes, err = db.CountReactionsToComment(1, comments[i].ID)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 		comments[i].Dislikes, err = db.CountReactionsToComment(0, comments[i].ID)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 	}

	// 	SendJSON(w, &comments)
}

func NewComment(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.Comments
	ReceiveJSON(r, &new)
	db.CreateComment(new)
}

func NewPost(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.PostDTO
	ReceiveJSON(r, &new)
	post := models.Post{}
	post.AuthorID = new.Author.ID
	post.Title = new.Title
	post.Content = new.Content
	ID, err := db.CreatePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := 0; i < len(new.CategoriesID); i++ {
		db.AssociateCategory(ID, new.CategoriesID[i])
	}
}

func NewReaction(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.Reactions
	ReceiveJSON(r, &new)
	db.CreateReaction(new)
}

func GetCategories(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	// 	a, err := db.ReturnCategories()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	result := models.Categories{AllCategories: a}
	// 	SendJSON(w, &result)
}
