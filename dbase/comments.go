package dbase

import (
	"fmt"
	"log"

	"github.com/Muslimah94/forum-back/models"
)

// CountComments ...
func (db *DataBase) CountComments(postID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Comments WHERE PostID = ?`, postID)
	if err != nil {
		fmt.Println("CountComments Query:", err)
		return 0, err
	}
	defer rows.Close()
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

// SelectComments ...
func (db *DataBase) SelectComments(postID int) ([]models.Comment, error) {

	rows, err := db.DB.Query(`SELECT ID, AuthorID, PostID, Content FROM Comments WHERE PostID = ?`, postID)
	if err != nil {
		log.Println("SelectComments Query:", err)
		return nil, err
	}
	defer rows.Close()
	var AllComments []models.Comment
	for rows.Next() {
		var p models.Comment
		err = rows.Scan(&p.ID, &p.AuthorID, &p.PostID, &p.Content)
		if err != nil {
			log.Println("SelectComments rows.Scan:", err)
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

// InsertComment ...
func (db *DataBase) InsertComment(new models.Comment) error {

	st, err := db.DB.Prepare(`INSERT INTO Comments (AuthorID, PostID, Content) VALUES (?,?,?)`)
	if err != nil {
		fmt.Println("InsertComment Prepare", err)
		return err
	}
	defer st.Close()
	_, err = st.Exec(new.AuthorID, new.PostID, new.Content)
	if err != nil {
		fmt.Println("InsertComment Exec", err)
		return err
	}
	return nil
}
