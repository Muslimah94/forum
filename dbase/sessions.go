package dbase

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	models "../models"
	uuid "github.com/satori/go.uuid"
)

// CreateSession ...
func (db *DataBase) CreateSession(new models.Session, tx *sql.Tx) (uuid.UUID, error) {
	fmt.Println("CreateSESSION")
	var err error
	new.UUID, err = uuid.NewV4()
	if err != nil {
		log.Println("CreateSession uuid.NewV4:", err)
		return new.UUID, err
	}
	new.ExpDate = time.Now().Add(time.Hour * 1).Unix()
	st, err := tx.Prepare(`INSERT INTO Sessions (UserID, UUID, ExpDate) VALUES (?,?,?)`)
	defer st.Close()
	if err != nil {
		fmt.Println("CreateSession Prepare", err)
		return new.UUID, err
	}
	_, err = st.Exec(new.UserID, new.UUID, new.ExpDate)
	if err != nil {
		fmt.Println("CreateSession Exec", err)
		return new.UUID, err
	}
	return new.UUID, nil
}

// SelectUserSession ...
func (db *DataBase) SelectUserSession(new models.Session) (models.Session, error) {

	var existing models.Session
	rows, err := db.DB.Query(`SELECT ID, UserID, UUID, ExpDate FROM Sessions WHERE UserID = ?`, new.UserID)
	defer rows.Close()
	if err != nil {
		fmt.Println("SelectUserSessions Query:", err)
		return existing, err
	}
	if rows.Next() {
		err = rows.Scan(&existing.ID, &existing.UserID, &existing.UUID, &existing.ExpDate)
		if err != nil {
			fmt.Println("SelectUserSessions rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserSessions rows:", err)
		return existing, err
	}
	return existing, nil
}

// CompareExpDate ...
func (db *DataBase) CompareExpDate(new models.Session) bool {
	return new.ExpDate < time.Now().Unix()
}

// UpdateSession ...
func (db *DataBase) UpdateSession(new models.Session, tx *sql.Tx) {
	stmt, err := tx.Prepare(`UPDATE Sessions SET UUID = ?, ExpDate = ? WHERE UserID = ?`)
	if err != nil {
		fmt.Println("UpdateSession Prepare[comment]", err)
		return err
	}
	_, err = stmt.Exec(new.Type, new.AuthorID, new.CommentID)
	if err != nil {
		fmt.Println("UpdateSession Exec[comment]", err)
		return err
	}
	defer stmt.Close()
}
