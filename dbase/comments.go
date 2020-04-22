package dbase

import (
	"fmt"

	models "../models"
)

func (db *DataBase) CountComments(postID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Comments WHERE PostID = ?`, postID)
	if err != nil {
		fmt.Println("CountComments Query:", err)
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountComments rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountComments rows:", err)
		return 0, err
	}
	return num, nil
}

func (db *DataBase) SelectComments(postID int) ([]models.Comments, error) {

	rows, err := db.DB.Query(`SELECT Comments.ID, AuthorID, PostID ,Content, Users.Nickname FROM Comments INNER JOIN
	Users ON Comments.AuthorID = Users.ID WHERE Comments.PostID = ?`, postID)
	if err != nil {
		fmt.Println("SelectComments Query:", err)
		return nil, err
	}
	var AllComments []models.Comments
	for rows.Next() {
		var p models.Comments
		err = rows.Scan(&p.ID, &p.AuthorID, &p.PostID, &p.Content, &p.AuthorNick)
		if err != nil {
			fmt.Println("SelectComments rows.Scan:", err)
			continue
		}
		AllComments = append(AllComments, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectComments rows:", err)
		return nil, err
	}
	return AllComments, nil
}

func (db *DataBase) CreateComment(new models.Comments) error {

	st, err := db.DB.Prepare(`INSERT INTO Comments (AuthorID, PostID, Content) VALUES (?,?,?)`)
	if err != nil {
		fmt.Println("CreateComment Prepare", err)
		return err
	}
	_, err = st.Exec(new.AuthorID, new.PostID, new.Content)
	if err != nil {
		fmt.Println("CreateComment Exec", err)
		return err
	}
	return nil
}

// func GetAllComments(db *sql.DB, w http.ResponseWriter, r *http.Request) {

// 	rows, err1 := db.Query(`SELECT * FROM Comments`)
// 	if err1 != nil {
// 		fmt.Println("GetAllComments db.Query ERROR:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	var AllComments []models.Comments
// 	for rows.Next() {
// 		var p models.Comments
// 		err2 := rows.Scan(&p.ID, &p.AuthorID, &p.PostID, &p.Content)
// 		if err2 != nil {
// 			fmt.Println("GetAllComments rows.Scan ERROR:", err2)
// 			continue
// 		}
// 		AllComments = append(AllComments, p)
// 	}
// 	if err3 := rows.Err(); err3 != nil {
// 		fmt.Println("GetAllComments rows ERROR:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, AllComments)
// }

// func GetCommentByID(db *sql.DB, w http.ResponseWriter, r *http.Request, commentID int) {
// 	var comment models.Comments
// 	rows := db.QueryRow(`SELECT * FROM Comments WHERE ID = $1`, commentID)
// 	err := rows.Scan(&comment.ID, &comment.AuthorID, &comment.PostID, &comment.Content)
// 	if err != nil {
// 		fmt.Println("GetCommentByID rows.Scan ERROR:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	handlers.SendJSON(w, comment)
// }

// func EditCommentByID(db *sql.DB, w http.ResponseWriter, r *http.Request, commentID int) {
// 	var new *models.Comments
// 	handlers.ReceiveJSON(r, &new)
// 	st, err2 := db.Prepare(`UPDATE Comments SET AuthorID = ?, PostID = ?, Content = ? where ID = ?`)
// 	if err2 != nil {
// 		fmt.Println("EditCommentsByID db.Prepare:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err3 := st.Exec(new.AuthorID, new.PostID, new.Content, commentID)
// 	if err3 != nil {
// 		fmt.Println("EditCommentsByID st.Exec:", err3)
// 		http.Error(w, err3.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func DeleteCommentByID(db *sql.DB, w http.ResponseWriter, r *http.Request, commentID int) {
// 	st, err1 := db.Prepare(`DELETE FROM Comments WHERE ID = ?`)
// 	if err1 != nil {
// 		fmt.Println("DeleteCommentByID db.Prepare:", err1)
// 		http.Error(w, err1.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	_, err2 := st.Exec(commentID)
// 	if err2 != nil {
// 		fmt.Println("DeleteCommentByID st.Exec:", err2)
// 		http.Error(w, err2.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
