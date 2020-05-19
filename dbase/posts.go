package dbase

import (
	"database/sql"
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
func (db *DataBase) SelectPostByID(postID int) (models.Post, error) {

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
func (db *DataBase) CreatePost(new models.Post, tx *sql.Tx) (int64, error) {
	var n int64
	d := time.Now().Unix()
	st, err := tx.Prepare(`INSERT INTO Posts (AuthorID, Title, Content, CreationDate) VALUES (?,?,?,?)`)
	defer st.Close()
	if err != nil {
		fmt.Println("CreatePost Prepare", err)
		return n, err
	}
	res, err := st.Exec(new.AuthorID, new.Title, new.Content, d)
	if err != nil {
		fmt.Println("CreatePost Exec", err)
		return n, err
	}
	n, err = res.LastInsertId()
	if err != nil {
		fmt.Println("InsertUser Exec", err)
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
