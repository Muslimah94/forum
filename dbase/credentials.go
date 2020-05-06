package dbase

import (
	"fmt"

	models "../models"
)

func (db *DataBase) CreateUserCredentials(new models.Credentials) error {
	st, err := db.DB.Prepare(`INSERT INTO Credentials (ID, Email, HashedPassword) VALUES (?,?,?)`)
	defer st.Close()
	if err != nil {
		fmt.Println("CreateUserCredentials Prepare", err)
		return err
	}
	_, err = st.Exec(new.ID, new.Email, new.HashedPassword)
	if err != nil {
		fmt.Println("CreateUserCredentials Exec", err)
		return err
	}
	return nil
}
