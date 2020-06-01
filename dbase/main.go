package dbase

import (
	"database/sql"
	"fmt"
	"os"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// DataBase ...
type DataBase struct {
	DB *sql.DB
}

// Create function creates or opens DB (if it's already exists)
func Create(DBname string) (*DataBase, error) {

	db, err := sql.Open("sqlite3", "./dbase/"+DBname)
	if err != nil {
		fmt.Println("Create sql.Open:", err)
		return nil, err
	}
	_, err = db.Exec(`PRAGMA foreign_keys = ON`)
	if err != nil {
		fmt.Println("Failed to set foreign keys in DB:", err)
		return nil, err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Users (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Nickname	TEXT NOT NULL UNIQUE,
		FirstName	TEXT,
		LastName	TEXT,
		Avatar	TEXT,
		RoleID	INTEGER NOT NULL,
		FOREIGN KEY(RoleID) REFERENCES Roles(ID) ON UPDATE CASCADE);
	CREATE TABLE IF NOT EXISTS Credentials (
		ID	INTEGER NOT NULL,
		Email	TEXT NOT NULL UNIQUE,
		HashedPassword	TEXT NOT NULL,
		PRIMARY KEY(ID),
		FOREIGN KEY(ID) REFERENCES Users(ID) ON DELETE CASCADE);
	CREATE TABLE IF NOT EXISTS Roles (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Name	TEXT NOT NULL UNIQUE);
	CREATE TABLE IF NOT EXISTS Categories (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Name	TEXT NOT NULL UNIQUE);
	CREATE TABLE IF NOT EXISTS Posts (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		AuthorID	INTEGER NOT NULL,
		Title	TEXT NOT NULL,
		Content	BLOB NOT NULL,
		CreationDate TEXT NOT NULL,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE);
	CREATE TABLE IF NOT EXISTS Comments (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		AuthorID	INTEGER NOT NULL,
		PostID	INTEGER NOT NULL,
		Content	BLOB NOT NULL,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE);
	CREATE TABLE IF NOT EXISTS Reactions (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		Type	INTEGER NOT NULL,
		AuthorID	INTEGER NOT NULL,
		PostID	INTEGER,
		CommentID	INTEGER,
		FOREIGN KEY(CommentID) REFERENCES Comments(ID) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE,
		FOREIGN KEY(AuthorID) REFERENCES Users(ID) ON DELETE CASCADE);
	CREATE TABLE IF NOT EXISTS Sessions (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		UserID	INTEGER NOT NULL UNIQUE,
		UUID	TEXT NOT NULL UNIQUE,
		ExpDate	TEXT NOT NULL,
		FOREIGN KEY(UserID) REFERENCES Users(ID) ON DELETE CASCADE);
	CREATE TABLE IF NOT EXISTS PostCats (
		ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		PostID	INTEGER NOT NULL,
		CategoryID	INTEGER NOT NULL,
		FOREIGN KEY(CategoryID) REFERENCES Categories(ID) ON DELETE CASCADE,
		FOREIGN KEY(PostID) REFERENCES Posts(ID) ON DELETE CASCADE);`)
	if err != nil {
		fmt.Println("Failed to create tables:", err)
		return nil, err
	}
	if !Exists("dbase/forumDB") {
		_, err = db.Exec(`
		INSERT INTO "main"."Roles"
			("Name")
		VALUES
			("admin"),
			("moderator"),
			("user");
		INSERT INTO "main"."Categories"
			("Name")
		VALUES
			("Web & Mobile dev"),
			("System dev"),
			("Graphics"),
			("Algorithms");`)
		if err != nil {
			fmt.Println("Failed to insert roles and categories:", err)
			return nil, err
		}
	}
	database := DataBase{DB: db}
	return &database, nil
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
