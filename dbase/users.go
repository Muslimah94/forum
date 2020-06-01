package dbase

import (
	"database/sql"
	"fmt"

	"github.com/Muslimah94/forum/models"
)

// SelectUsers ...
func (db *DataBase) SelectUsers() ([]models.User, error) {
	rows, err := db.DB.Query(`SELECT ID, Nickname, RoleID FROM Users`)
	if err != nil {
		fmt.Println("SelectUsers db.Query ERROR:", err)
		return nil, err
	}
	defer rows.Close()
	var AllUsers []models.User
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Nickname, &u.RoleID)
		if err != nil {
			fmt.Println("SelectUsers rows.Scan ERROR:", err)
			continue
		}
		AllUsers = append(AllUsers, u)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUsers rows ERROR:", err)
		return nil, err
	}
	return AllUsers, nil
}

// SelectUserByID ...
func (db *DataBase) SelectUserByID(userID int) (models.User, error) {
	var u models.User
	rows, err := db.DB.Query(`SELECT ID, Nickname, RoleID FROM Users WHERE ID = ?`, userID)
	if err != nil {
		fmt.Println("SelectUserByID db.Query ERROR:", err)
		return u, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&u.ID, &u.Nickname, &u.RoleID)
		if err != nil {
			fmt.Println("SelectUserByID rows.Scan ERROR:", err)
			return u, err
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserByID rows ERROR:", err)
		return u, err
	}
	return u, nil
}

// InsertUser ...
func (db *DataBase) InsertUser(new models.User, tx *sql.Tx) (int64, error) {
	var n int64
	st, err := tx.Prepare(`INSERT INTO Users (Nickname, RoleID) VALUES (?,?)`)
	if err != nil {
		fmt.Println("InsertUser Prepare", err)
		return n, err
	}
	defer st.Close()
	res, err := st.Exec(new.Nickname, new.RoleID)
	if err != nil {
		fmt.Println("InsertUser Exec", err)
		return n, err
	}
	n, err = res.LastInsertId()
	if err != nil {
		fmt.Println("InsertUser Exec", err)
		return n, err
	}
	return n, nil
}

// SelectUserBySession ...
func (db *DataBase) SelectUserIDBySession(UUID string) (int, error) {
	var id int
	rows, err := db.DB.Query(`SELECT UserID FROM Sessions WHERE UUID = ?`, UUID)
	if err != nil {
		fmt.Println("SelectUserBySession db.Query ERROR:", err)
		return id, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("SelectUserBySession rows.Scan ERROR:", err)
			return id, err
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserBySession rows ERROR:", err)
		return id, err
	}
	return id, nil
}
