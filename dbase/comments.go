package dbase

import (
	"database/sql"
	"fmt"
	"net/http"

	models "../models"
)

func GetAllComments(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	rows, err1 := db.Query(`SELECT * FROM Comments`)
	if err1 != nil {
		fmt.Println("GetAllComments db.Query ERROR:", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	var AllComments []models.Comments
	for rows.Next() {
		var p models.Comments
		err2 := rows.Scan(&p.ID, &p.AuthorID, &p.PostID, &p.Content)
		if err2 != nil {
			fmt.Println("GetAllComments rows.Scan ERROR:", err2)
			continue
		}
		AllComments = append(AllComments, p)
	}
	if err3 := rows.Err(); err3 != nil {
		fmt.Println("GetAllComments rows ERROR:", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
	SendJSON(w, AllComments)
}

func AddNewComment(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var post *models.Comments
	ReceiveJSON(r, &post)
	st, err1 := db.Prepare(`INSERT INTO Comments (AuthorID,PostID,Content) VALUES (?,?,?)`)
	if err1 != nil {
		fmt.Println("AddNewComments db.Prepare", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	_, err2 := st.Exec(post.AuthorID, post.PostID, post.Content)
	if err2 != nil {
		fmt.Println("AddNewComments st.Exec", err2)
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCommentByID(db *sql.DB, w http.ResponseWriter, r *http.Request, commentID int) {
	var comment models.Comments
	rows := db.QueryRow(`SELECT * FROM Comments WHERE ID = $1`, commentID)
	err := rows.Scan(&comment.ID, &comment.AuthorID, &comment.PostID, &comment.Content)
	if err != nil {
		fmt.Println("GetCommentByID rows.Scan ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJSON(w, comment)
}

func EditCommentByID(db *sql.DB, w http.ResponseWriter, r *http.Request, commentID int) {
	var new *models.Comments
	ReceiveJSON(r, &new)
	st, err2 := db.Prepare(`UPDATE Comments SET AuthorID = ?, PostID = ?, Content = ? where ID = ?`)
	if err2 != nil {
		fmt.Println("EditCommentsByID db.Prepare:", err2)
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	_, err3 := st.Exec(new.AuthorID, new.PostID, new.Content, commentID)
	if err3 != nil {
		fmt.Println("EditCommentsByID st.Exec:", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteCommentByID(db *sql.DB, w http.ResponseWriter, r *http.Request, commentID int) {
	st, err1 := db.Prepare(`DELETE FROM Comments WHERE ID = ?`)
	if err1 != nil {
		fmt.Println("DeleteCommentByID db.Prepare:", err1)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	_, err2 := st.Exec(commentID)
	if err2 != nil {
		fmt.Println("DeleteCommentByID st.Exec:", err2)
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}
