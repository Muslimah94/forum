package dbase

import (
	"database/sql"
	"fmt"

	models "../models"
)

func (db *DataBase) CreateUserCredentials(new models.Credentials, tx *sql.Tx) error {
	fmt.Println("CreateUserCREDENTIALS")
	st, err := tx.Prepare(`INSERT INTO Credentials (ID, Email, HashedPassword) VALUES (?,?,?)`)
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

func (db *DataBase) SelectUserCredentials(new models.Credentials) (models.Credentials, error) {

	var existing models.Credentials
	rows, err := db.DB.Query(`SELECT ID, Email, HashedPassword FROM Credentials WHERE Email = ?`, new.Email)
	defer rows.Close()
	if err != nil {
		fmt.Println("SelectUserCredentials Query:", err)
		return existing, err
	}
	if rows.Next() {
		err = rows.Scan(&existing.ID, &existing.Email, &existing.HashedPassword)
		if err != nil {
			fmt.Println("SelectUserCredentials rows.Scan:", err)
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println("SelectUserCredentials rows:", err)
		return existing, err
	}
	return existing, nil
}
