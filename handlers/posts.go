package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Muslimah94/forum/dbase"
	"github.com/Muslimah94/forum/models"
)

// GetAllPosts ...
func GetAllPosts(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	//-------FLAGS FOR FILTER----------------------------------------
	all := false
	liked := false
	created := false
	//-------GETTING PARAMETERS FOR FILTER---------------------------
	l, ok := r.URL.Query()["liked"]
	if !ok || len(l[0]) < 1 {
		log.Println("GetAllPosts: Url Param 'liked' is missing")
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	c, ok := r.URL.Query()["created"]
	if !ok || len(c[0]) < 1 {
		log.Println("GetAllPosts: Url Param 'created' is missing")
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	//--------FILTERING ---------------------------------------------
	UserID, err := GetUserIDBySession(db, r)
	if l[0] == "0" && c[0] == "0" {
		all = true
	} else if l[0] == "1" {
		liked = true
	} else if c[0] == "1" {
		created = true
	}
	if _, err := strconv.Atoi(l[0]); err == nil && l[0] != "1" && l[0] != "0" {
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	if _, err := strconv.Atoi(l[0]); err == nil && c[0] != "1" && c[0] != "0" {
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	if (err != nil || UserID == 0) && liked {
		SendJSON(w, models.Error{
			Status:      "Unathorized",
			Description: "Please authorize to see your liked posts",
		})
		return
	}
	if (err != nil || UserID == 0) && created {
		SendJSON(w, models.Error{
			Status:      "Unathorized",
			Description: "Please authorize to see your created posts",
		})
		return
	}
	//-------ENTITY--------------------------------------------------
	posts := []models.Post{}
	if all {
		posts, err = db.SelectPosts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if liked {
		IDs, err := db.SelectLikedPostsIDs(UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i := 0; i < len(IDs); i++ {
			p, err := db.SelectPostByID(IDs[i])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			posts = append(posts, p)
		}
	} else if created {
		posts, err = db.SelectCreatedPosts(UserID)
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
	//-------------DTO-----------------------------------------------
	DTOs := []models.PostDTO{}
	for i := 0; i < len(posts); i++ {
		pDTO := models.PostDTO{}
		// post id ----------
		pDTO.ID = posts[i].ID
		for _, v := range users {
			if v.ID == posts[i].AuthorID {
				a := models.AuthorDTO{}
				a.ID = v.ID
				a.Nickname = v.Nickname
				pDTO.Author = a
			}
		}
		// post title -----------
		pDTO.Title = posts[i].Title
		// post content ---------
		pDTO.Content = posts[i].Content
		// post categories ------
		ar := []string{}
		for j := 0; j < len(pc); j++ {
			if posts[i].ID == pc[j].PostID {
				ar = append(ar, pc[j].CategoryName)
			}
		}
		pDTO.Categories = ar
		// number of likes ------
		pDTO.Likes, err = db.CountReactionsToPost(1, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// number of dislikes ---
		pDTO.Dislikes, err = db.CountReactionsToPost(0, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// number of comments ---
		pDTO.Comments, err = db.CountComments(posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// post creation date ---
		pDTO.CreationDate = posts[i].CreationDate

		// reaction of session user to post
		if UserID != 0 {
			reaction, err := db.SelectReaction(models.Reaction{
				AuthorID: UserID,
				PostID:   posts[i].ID,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if reaction.AuthorID == 0 {
				pDTO.UserReaction = -1
			} else {
				pDTO.UserReaction = reaction.Type
			}
		} else {
			pDTO.UserReaction = -1
		}
		DTOs = append(DTOs, pDTO)
	}
	SendJSON(w, &DTOs)
}

// GetPostByID ...
func GetPostByID(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	p, ok := r.URL.Query()["id"]
	if !ok || len(p[0]) < 1 {
		log.Println("GetPostByID: Url Param 'id' is missing")
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(p[0])
	if err != nil {
		log.Println("GetPostByID Atoi:", err)
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	//----------ENTITY---------------------------------------------------------------
	post, err := db.SelectPostByID(postID)
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
	postDTO.CreationDate = post.CreationDate

	id, err := GetUserIDBySession(db, r)
	if err != nil {
		postDTO.UserReaction = -1
		SendJSON(w, &postDTO)
		return
	}
	reaction, err := db.SelectReaction(models.Reaction{
		AuthorID: id,
		PostID:   post.ID,
	})
	if reaction.AuthorID == 0 || err != nil {
		postDTO.UserReaction = -1
	} else {
		postDTO.UserReaction = reaction.Type
	}
	SendJSON(w, &postDTO)
}

// NewPost ...
func NewPost(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	var new models.PostDTO
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := GetUserIDBySession(db, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post := models.Post{}
	post.AuthorID = id
	post.Title = new.Title
	post.Content = new.Content
	//-------STARTING TRANSACTION--------
	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Cannot start transaction")
		return
	}
	ID, err := db.InsertPost(post, tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	for i := 0; i < len(new.CategoriesID); i++ {
		err = db.AssociateCategory(int(ID), new.CategoriesID[i], tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

// GetPostsByCategory ...
func GetPostsByCategory(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	cat, ok := r.URL.Query()["category"]
	if !ok || len(cat[0]) < 1 {
		log.Println("GetPostsByCategory: Url Param 'category' is missing")
		http.Error(w, "Bad request, please try again", http.StatusBadRequest)
		return
	}
	PC, err := db.SelectPostIDsByCategory(cat[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts := []models.Post{}
	for i := 0; i < len(PC); i++ {
		p, err := db.SelectPostByID(PC[i].PostID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
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

	// Copying again -_-
	DTOs := []models.PostDTO{}
	for i := 0; i < len(posts); i++ {
		pDTO := models.PostDTO{}
		// post id ----------
		pDTO.ID = posts[i].ID
		for _, v := range users {
			if v.ID == posts[i].AuthorID {
				a := models.AuthorDTO{}
				a.ID = v.ID
				a.Nickname = v.Nickname
				pDTO.Author = a
			}
		}
		// post title -----------
		pDTO.Title = posts[i].Title
		// post content ---------
		pDTO.Content = posts[i].Content
		// post categories ------
		ar := []string{}
		for j := 0; j < len(pc); j++ {
			if posts[i].ID == pc[j].PostID {
				ar = append(ar, pc[j].CategoryName)
			}
		}
		pDTO.Categories = ar
		// number of likes ------
		pDTO.Likes, err = db.CountReactionsToPost(1, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// number of dislikes ---
		pDTO.Dislikes, err = db.CountReactionsToPost(0, posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// number of comments ---
		pDTO.Comments, err = db.CountComments(posts[i].ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// post creation date ---
		pDTO.CreationDate = posts[i].CreationDate

		// reaction of session user to post
		UserID, err := GetUserIDBySession(db, r)
		if err != nil || UserID == 0 {
			pDTO.UserReaction = -1
		} else {
			reaction, err := db.SelectReaction(models.Reaction{
				AuthorID: UserID,
				PostID:   posts[i].ID,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if reaction.AuthorID == 0 {
				pDTO.UserReaction = -1
			} else {
				pDTO.UserReaction = reaction.Type
			}
		}
		DTOs = append(DTOs, pDTO)
	}
	SendJSON(w, &DTOs)

}
