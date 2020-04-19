package dbase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Create function creates or opens DB (if it's already exists)
func Create(DBname string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./"+DBname)
	if err != nil {
		fmt.Println("Create sql.Open:", err)
		return nil, err
	}
	_, err1 := db.Exec(`PRAGMA foreign_keys = ON`)
	if err1 != nil {
		fmt.Println("Failed to set foreign keys:", err1)
		return nil, err1
	}
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Email	TEXT NOT NULL UNIQUE,
		Nickname	TEXT NOT NULL UNIQUE,
		Password	TEXT NOT NULL,
		RoleID	INTEGER NOT NULL,
		FOREIGN KEY (RoleID) REFERENCES Roles(ID))`)
	if err2 != nil {
		fmt.Println("Failed to create Users table:", err2)
		return nil, err2
	}
	_, err3 := db.Exec(`CREATE TABLE IF NOT EXISTS Roles (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Name	TEXT NOT NULL)`)
	if err3 != nil {
		fmt.Println("Failed to create Roles table:", err3)
		return nil, err3
	}
	_, err4 := db.Exec(`CREATE TABLE IF NOT EXISTS Categories (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Name	TEXT NOT NULL UNIQUE)`)
	if err4 != nil {
		fmt.Println("Failed to create Categories table:", err4)
		return nil, err4
	}
	_, err5 := db.Exec(`CREATE TABLE IF NOT EXISTS Posts (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		AuthorID	INTEGER NOT NULL,
		Title	TEXT NOT NULL,
		Content	TEXT NOT NULL,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE)`)
	if err5 != nil {
		fmt.Println("Failed to create Posts table:", err5)
		return nil, err5
	}
	_, err6 := db.Exec(`CREATE TABLE IF NOT EXISTS Comments (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		AuthorID	INTEGER NOT NULL,
		PostID	INTEGER NOT NULL,
		Content	INTEGER NOT NULL,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE)`)
	if err6 != nil {
		fmt.Println("Failed to create Comments table:", err6)
		return nil, err6
	}
	_, err7 := db.Exec(`CREATE TABLE IF NOT EXISTS Reactions (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Type	INTEGER NOT NULL,
		AuthorID	INTEGER NOT NULL,
		PostID	INTEGER,
		CommentID	INTEGER,
		FOREIGN KEY(CommentID) REFERENCES Comments(ID) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE)`)
	if err7 != nil {
		fmt.Println("Failed to create Reactions table:", err7)
		return nil, err7
	}
	_, err8 := db.Exec(`CREATE TABLE IF NOT EXISTS Sessions (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		UserID	INTEGER NOT NULL,
		UUID	TEXT NOT NULL,
		ExpDate	TEXT NOT NULL,
		FOREIGN KEY(UserID) REFERENCES Users(ID) ON DELETE CASCADE)`)
	if err8 != nil {
		fmt.Println("Failed to create Sessions table:", err8)
		return nil, err8
	}
	_, err9 := db.Exec(`CREATE TABLE IF NOT EXISTS PostsCategories (
		PostID	INTEGER NOT NULL,
		CategoryID	INTEGER NOT NULL,
		FOREIGN KEY(CategoryID) REFERENCES Categories(ID) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE)`)
	if err9 != nil {
		fmt.Println("Failed to create PostsCategories table:", err9)
		return nil, err9
	}
	return db, nil
}

// SendJSON function marshals and sends given data to response writer
func SendJSON(v interface{}) {
	var w http.ResponseWriter
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println("SendJson json.Marshal ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// ReceiveJSON function decodes data from request
func ReceiveJSON(r *http.Request, v interface{}) {
	var w http.ResponseWriter
	err1 := json.NewDecoder(r.Body).Decode(v)
	if err1 != nil {
		fmt.Println("ReceiveJSON: Failed to Decode", err1)
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
}
