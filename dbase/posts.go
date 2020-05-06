package dbase

import (
	"fmt"
	"time"

	models "../models"
)

// SelectPosts ...
func (db *DataBase) SelectPosts() ([]models.Post, error) {

	rows, err := db.DB.Query(`SELECT * FROM Posts`)
	defer rows.Close()
	if err != nil {
		fmt.Println("SelectPosts Query:", err)
		return nil, err
	}
	var AllPosts []models.Post
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreationDate)
		if err != nil {
			fmt.Println("SelectPosts rows.Scan:", err)
			continue
		}
		AllPosts = append(AllPosts, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectPosts rows:", err)
		return nil, err
	}
	return AllPosts, nil

}

// SelectPost ...
func (db *DataBase) SelectPost(postID int) (models.Post, error) {

	var p models.Post
	rows := db.DB.QueryRow(`SELECT * FROM Posts WHERE ID = ? `, postID)
	err := rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreationDate)
	if err != nil {
		fmt.Println("SelectPost:", err)
		return p, err
	}
	return p, nil
}

// CreatePost ...
func (db *DataBase) CreatePost(new models.Post) (int, error) {
	n := 0
	d := time.Now().Unix()
	st, err := db.DB.Prepare(`INSERT INTO Posts (AuthorID, Title, Content, CreationDate) VALUES (?,?,?,?)`)
	defer st.Close()
	if err != nil {
		fmt.Println("CreatePost Prepare", err)
		return n, err
	}
	_, err = st.Exec(new.AuthorID, new.Title, new.Content, d)
	if err != nil {
		fmt.Println("CreatePost Exec", err)
		return n, err
	}
	n, err = db.ReturnLastPostID()
	if err != nil {
		fmt.Println("CreatePost Exec", err)
		return n, err
	}
	return n, nil
}

// ReturnLastPostID ...
func (db *DataBase) ReturnLastPostID() (int, error) {
	n := 0
	rows, err := db.DB.Query(`SELECT ID FROM Posts ORDER BY ID DESC LIMIT 1`)
	defer rows.Close()
	if err != nil {
		fmt.Println("ReturnLastPostID Query:", err)
		return n, err
	}
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			fmt.Println("ReturnLastPostID rows.Scan:", err)
			continue
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("ReturnLastPostID rows:", err)
		return n, err
	}
	return n, nil
}

// // AddNewPost ...
// func AddNewPost(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	var post *models.Posts
// 	ReceiveJSON(r, &post)

// 	post.CreationDate = time.Now().Unix()

// 	st, err1 := db.Prepare(`INSERT INTO Posts (AuthorID,Title,Content, CreationDate) VALUES (?,?,?,?)`)
// 	if err1 != nil {
// 		fmt.Println("AddNewPost db.Prepare", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(post.AuthorID, post.Title, post.Content, post.CreationDate)
// 	if err3 != nil {
// 		fmt.Println("AddNewPost st.Exec", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
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
