package dbase

import (
	"fmt"
	"log"
	"time"

	models "../models"
	uuid "github.com/satori/go.uuid"
)

// CreateSession ...
func (db *DataBase) CreateSession(new models.Session) (uuid.UUID, error) {
	fmt.Println("CreateSESSION")
	var err error
	new.UUID, err = uuid.NewV4()
	if err != nil {
		log.Println("CreateSession uuid.NewV4:", err)
		return new.UUID, err
	}
	new.ExpDate = time.Now().Add(time.Hour * 1).Unix()
	st, err := db.DB.Prepare(`INSERT INTO Sessions (UserID, UUID, ExpDate) VALUES (?,?,?)`)
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
