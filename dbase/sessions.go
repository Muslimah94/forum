package dbase

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	models "../models"
	uuid "github.com/satori/go.uuid"
)

// InsertSession ...
func (db *DataBase) InsertSession(new models.Session, tx *sql.Tx) (uuid.UUID, error) {
	var err error
	new.UUID, err = uuid.NewV4()
	if err != nil {
		log.Println("InsertSession uuid.NewV4:", err)
		return new.UUID, err
	}
	new.ExpDate = time.Now().Add(time.Hour * 1).Unix()
	st, err := tx.Prepare(`INSERT INTO Sessions (UserID, UUID, ExpDate) VALUES (?,?,?)`)
	if err != nil {
		fmt.Println("InsertSession Prepare", err)
		return new.UUID, err
	}
	defer st.Close()
	_, err = st.Exec(new.UserID, new.UUID, new.ExpDate)
	if err != nil {
		fmt.Println("InsertSession Exec", err)
		return new.UUID, err
	}
	return new.UUID, err
}

// SelectUserSession ...
func (db *DataBase) SelectUserSession(new models.Session) (models.Session, error) {

	var existing models.Session
	rows, err := db.DB.Query(`SELECT ID, UserID, UUID, ExpDate FROM Sessions WHERE UserID = ?`, new.UserID)
	if err != nil {
		fmt.Println("SelectUserSession Query:", err)
		return existing, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&existing.ID, &existing.UserID, &existing.UUID, &existing.ExpDate)
		if err != nil {
			fmt.Println("SelectUserSession rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserSession rows:", err)
		return existing, err
	}
	return existing, nil
}

// UpdateSession ...
func (db *DataBase) UpdateSession(new models.Session, tx *sql.Tx) error {
	stmt, err := tx.Prepare(`UPDATE Sessions SET UUID = ?, ExpDate = ? WHERE UserID = ?`)
	if err != nil {
		fmt.Println("UpdateSession Prepare", err)
		return err
	}
	_, err = stmt.Exec(new.UUID, new.ExpDate, new.UserID)
	if err != nil {
		fmt.Println("UpdateSession Exec", err)
		return err
	}
	defer stmt.Close()
	return err
}

// UpdateSessionDate ...
func (db *DataBase) UpdateSessionDate(new models.Session, tx *sql.Tx) error {
	stmt, err := tx.Prepare(`UPDATE Sessions SET ExpDate = ? WHERE UserID = ?`)
	if err != nil {
		fmt.Println("UpdateSessionDate Prepare", err)
		return err
	}
	_, err = stmt.Exec(new.ExpDate, new.UserID)
	if err != nil {
		fmt.Println("UpdateSessionDate Exec", err)
		return err
	}
	defer stmt.Close()
	return err
}

// CompareExpDate ...
func (db *DataBase) CompareExpDate(new models.Session) bool {
	return new.ExpDate > time.Now().Unix()
}

// SelectSession ..
func (db *DataBase) SelectSession(UUID string) (models.Session, error) {
	var existing models.Session
	rows, err := db.DB.Query(`SELECT ID, UserID, UUID, ExpDate FROM Sessions WHERE UUID = ?`, UUID)
	if err != nil {
		fmt.Println("SelectUserSession Query:", err)
		return existing, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&existing.ID, &existing.UserID, &existing.UUID, &existing.ExpDate)
		if err != nil {
			fmt.Println("SelectUserSession rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserSession rows:", err)
		return existing, err
	}
	return existing, nil
}

// DeleteSession ..
func (db *DataBase) DeleteSession(UUID string) error {

	stmt, err := db.DB.Prepare(`DELETE FROM Sessions WHERE UUID = ?`)
	if err != nil {
		fmt.Println("DeleteUserSession Prepare:", err)
		return err
	}
	_, err = stmt.Exec(UUID)
	if err != nil {
		fmt.Println("DeleteUserSession Exec", err)
		return err
	}
	defer stmt.Close()
	return err
}
