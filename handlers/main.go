package handlers

import (
	"log"
	"net/http"
	"strconv"

	dbase "../dbase"
	models "../models"
)

func GetAllPosts(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	//-------ENTITY---------------------------------------------------------
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//---------------------DTO---------------------------
	DTOs := []models.PostDTO{}
	for i := 0; i < len(posts); i++ {
		pDTO := models.PostDTO{}
		pDTO.ID = posts[i].ID
		for _, v := range users {
			if v.ID == posts[i].AuthorID {
				a := models.AuthorDTO{}
				a.ID = v.ID
				a.Nickname = v.Nickname
				pDTO.Author = a
			}
		}
		pDTO.Title = posts[i].Title
		pDTO.Content = posts[i].Content
		ar := []string{}
		for j := 0; j < len(pc); j++ {
			if posts[i].ID == pc[j].PostID {
				ar = append(ar, pc[j].CategoryName)
			}
		}
		pDTO.Categories = ar
		pDTO.Likes, err = db.CountReactionsToPost(1, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pDTO.Dislikes, err = db.CountReactionsToPost(0, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pDTO.Comments, err = db.CountComments(posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pDTO.CreationDate = posts[i].CreationDate
		DTOs = append(DTOs, pDTO)
	}

	SendJSON(w, &DTOs)
}

func GetPostByID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	p, ok := r.URL.Query()["id"]
	if !ok || len(p[0]) < 1 {
		log.Println("GetPostByID: Url Param 'id' is missing")
		http.Error(w, "Internal Server Error, please try again later", http.StatusInternalServerError)
		return
	}
	postID, err := strconv.Atoi(p[0])
	if err != nil {
		log.Println("GetPostByID Atoi:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//----------ENTITY---------------------------------------------------------------
	post, err := db.SelectPost(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//-------DTO PREPARATION---------------------------------------------------------
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

func NewComment(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.CommentDTO
	ReceiveJSON(r, &new)
	//--------ENTITY---------------------------------------
	c := models.Comment{}
	c.AuthorID = new.Author.ID
	c.PostID = new.PostID
	c.Content = new.Content
	db.CreateComment(c)
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

func GetCategories(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	a, err := db.ReturnCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dto := models.CategoriesDTO{AllCategories: a}
	SendJSON(w, &dto)
}

func NewReaction(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.ReactionDTO
	ReceiveJSON(r, &new)
	//-------ENTITY---------------
	rea := models.Reaction{
		AuthorID:  new.AuthorID,
		Type:      new.Type,
		PostID:    new.PostID,
		CommentID: new.CommentID,
	}
	db.CreateReaction(rea)
}
