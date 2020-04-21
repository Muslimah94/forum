package dbase

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	models "../models"
)

// GetAllPosts ...
func GetAllPosts(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	rows, err1 := db.Query(`SELECT Posts.ID, Posts.Title, Content, CreationDate, Users.Nickname FROM Posts INNER JOIN
	Users ON Posts.AuthorID = Users.ID`)
	if err1 != nil {
		fmt.Println("GetAllPosts db.Query ERROR:", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	var AllPosts []models.Posts
	for rows.Next() {
		var p models.Posts
		err2 := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreationDate, &p.AuthorNick)
		if err2 != nil {
			fmt.Println("GetAllPosts rows.Scan ERROR:", err2)
			continue
		}
		AllPosts = append(AllPosts, p)
	}
	if err3 := rows.Err(); err3 != nil {
		fmt.Println("GetAllPosts rows ERROR:", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	// --------------------------------------------------------------
	rows2, err1 := db.Query(`SELECT PostsCategories.PostID, CategoryID, Categories.Name FROM PostsCategories INNER JOIN
	Categories ON PostsCategories.CategoryID = Categories.ID`)
	if err1 != nil {
		fmt.Println("GetAllPostsCat db.Query ERROR:", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	var pc []models.PostsCategories
	for rows2.Next() {
		var p models.PostsCategories
		err2 := rows2.Scan(&p.PostID, &p.CategoryID, &p.CategoryName)
		if err2 != nil {
			fmt.Println("GetAllPostsCat rows.Scan ERROR:", err2)
			continue
		}
		pc = append(pc, p)
	}
	if err3 := rows.Err(); err3 != nil {
		fmt.Println("GetAllPostsCat rows ERROR:", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range AllPosts {
		fmt.Println("v.Cat:", v.Categories)
		ar := []string{}
		for _, g := range pc {
			if v.ID == g.PostID {
				ar = append(ar, g.CategoryName)
			}
		}
		v.Categories = ar
		fmt.Println("v.Cat after:", v.Categories)
	}
	SendJSON(w, AllPosts)
}

// AddNewPost ...
func AddNewPost(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var post *models.Posts
	ReceiveJSON(r, &post)

	post.CreationDate = time.Now().Unix()

	st, err1 := db.Prepare(`INSERT INTO Posts (AuthorID,Title,Content, CreationDate) VALUES (?,?,?,?)`)
	if err1 != nil {
		fmt.Println("AddNewPost db.Prepare", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	_, err3 := st.Exec(post.AuthorID, post.Title, post.Content, post.CreationDate)
	if err3 != nil {
		fmt.Println("AddNewPost st.Exec", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}

// // GetPostByID ...
// func GetPostByID(db *sql.DB, w http.ResponseWriter, r *http.Request, postID int) {

// 	var post models.Posts
// 	rows := db.QueryRow(`SELECT * FROM Posts WHERE ID = $1`, postID)
// 	err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreationDate)
// 	if err != nil {
// 		fmt.Println("GetPostByID rows.Scan ERROR:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var user models.Users
// 	rows2 := db.QueryRow(`SELECT * FROM Users WHERE ID = $1`, post.AuthorID)
// 	err2 := rows2.Scan(&user.ID, &user.Email, &user.Nickname, &user.Password, &user.RoleID)
// 	post.AuthorNick = user.Nickname
// 	if err2 != nil {
// 		fmt.Println("GetAllPosts2 rows.Scan ERROR:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	SendJSON(w, post)
// }

// // EditPostByID ...
// func EditPostByID(db *sql.DB, w http.ResponseWriter, r *http.Request, postID int) {
// 	var new *models.Posts
// 	ReceiveJSON(r, &new)
// 	new.CreationDate = time.Now().Unix()

// 	var user models.Users
// 	rows := db.QueryRow(`SELECT * FROM Users WHERE ID = $1`, new.AuthorID)
// 	err := rows.Scan(&user.ID, &user.Email, &user.Nickname, &user.Password, &user.RoleID)
// 	new.AuthorNick = user.Nickname
// 	if err != nil {
// 		fmt.Println("GetAllPosts2 rows.Scan ERROR:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	st, err2 := db.Prepare(`UPDATE Posts SET AuthorID = ?, Title = ?, Content = ?, CreationDate = ? WHERE ID = ?`)
// 	if err2 != nil {
// 		fmt.Println("EditPostByID db.Prepare:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(new.AuthorID, new.Title, new.Content, new.CreationDate, postID)
// 	if err3 != nil {
// 		fmt.Println("EditUserByID st.Exec:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // DeletePostByID ...
// func DeletePostByID(db *sql.DB, w http.ResponseWriter, r *http.Request, postID int) {
// 	st, err1 := db.Prepare(`DELETE FROM Posts WHERE ID = ?`)
// 	if err1 != nil {
// 		fmt.Println("DeletePostByID db.Prepare:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err2 := st.Exec(postID)
// 	if err2 != nil {
// 		fmt.Println("DeletePostByID st.Exec:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// // GetPostsByCategoryID ...
// func GetPostsByCategoryID(db *sql.DB, w http.ResponseWriter, r *http.Request, categoryID int) {

// 	rows, err1 := db.Query(`SELECT * FROM PostsCategories WHERE CategoryID = ?`, categoryID)
// 	if err1 != nil {
// 		fmt.Println("GetPostsByCategoryID db.Query ERROR:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	var PostCategory []models.PostsCategories
// 	for rows.Next() {
// 		var p models.PostsCategories
// 		err2 := rows.Scan(&p.PostID, &p.CategoryID)
// 		if err2 != nil {
// 			fmt.Println("GetPostsByCategoryID rows.Scan ERROR:", err2)
// 			continue
// 		}
// 		PostCategory = append(PostCategory, p)
// 	}
// 	if err3 := rows.Err(); err3 != nil {
// 		fmt.Println("GetPostsByCategoryID rows ERROR:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var Posts []models.Posts
// 	for _, v := range PostCategory {
// 		var post models.Posts
// 		rows := db.QueryRow(`SELECT * FROM Posts WHERE ID = $1`, v.PostID)
// 		err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreationDate)
// 		if err != nil {
// 			fmt.Println("GetPostByID rows.Scan ERROR:", err)
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		Posts = append(Posts, post)
// 	}

// 	var user models.Users
// 	for _, v := range Posts {
// 		rows := db.QueryRow(`SELECT * FROM Users WHERE ID = $1`, v.AuthorID)
// 		err := rows.Scan(&user.ID, &user.Email, &user.Nickname, &user.Password, &user.RoleID)
// 		v.AuthorNick = user.Nickname
// 		if err != nil {
// 			fmt.Println("GetPosts2 rows.Scan ERROR:", err)
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// 	SendJSON(w, Posts)
// }
