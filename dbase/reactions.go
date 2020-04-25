package dbase

import (
	"fmt"
)

func (db *DataBase) CountReactionsToPost(t int, postID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND PostID = ?`, t, postID)
	defer rows.Close()
	if err != nil {
		fmt.Println("CountReactionsToPost Query:", err)
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountReactionsToPost rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountReactionsToPost rows:", err)
		return 0, err
	}
	return num, nil
}

func (db *DataBase) CountReactionsToComment(t int, commentID int) (int, error) {
	num := 0
	rows, err := db.DB.Query(`SELECT COUNT(*) FROM Reactions WHERE Type = ? AND CommentID = ?`, t, commentID)
	defer rows.Close()
	if err != nil {
		fmt.Println("CountReactionsToComment Query:", err)
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("CountReactionsToComment rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("CountReactionsToComment rows:", err)
		return 0, err
	}
	return num, nil
}
