package dbase

import (
	"database/sql"
	"fmt"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// DataBase ...
type DataBase struct {
	DB *sql.DB
}

// Create function creates or opens DB (if it's already exists)
func Create(DBname string) (*DataBase, error) {

	db, err := sql.Open("sqlite3", "./"+DBname)
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
		UserID	INTEGER NOT NULL,
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
	database := DataBase{DB: db}
	return &database, nil
}

// _, err = db.Exec(`INSERT INTO Roles (Name)
// 	VALUES
// 		("Admin"),
// 		("Moderator"),
// 		("User"),
// 		("Guest");
// 	INSERT INTO Users (Nickname, FirstName, LastName, RoleID)
// 	VALUES
// 		("Pirozhok", "Serik", "Serik", 1),
// 		("Alibek-tse", "Alibek", "Tokanov", 2),
// 		("Muslimah94", "Maral", "Tokanova", 2),
// 		("Koba", "Kobylan", "Kobylan", 3),
// 		("Alanapapa", "Berik", "Berik", 3);
// 	INSERT INTO Posts (AuthorID, Title, Content)
// 	VALUES
// 		(1, "Post number 1", "Serik Serik Serik Serik Serik Serik Serik"),
// 		(2, "Post number 2", "Alibek Alibek Alibek Alibek Alibek Alibek Alibek");
// 	INSERT INTO Reactions (Type, AuthorID, PostID)
// 	VALUES
// 		(0, 1, 1),
// 		(0, 2, 1),
// 		(0, 3, 1),
// 		(0, 4, 1),
// 		(1, 5, 1),
// 		(1, 1, 2),
// 		(1, 2, 2),
// 		(1, 3, 2),
// 		(1, 4, 2),
// 		(0, 5, 2);
// 	INSERT INTO Categories (Name)
// 	VALUES
// 		("Web and Mobile Development"),
// 		("Algorithms"),
// 		("Graphics"),
// 		("IU");
// 	INSERT INTO PostCats (PostID, CategoryID)
// 	VALUES
// 		(1, 1),
// 		(1, 2),
// 		(2, 3),
// 		(2, 4);
// 	INSERT INTO Comments (AuthorID, PostID, Content)
// 	VALUES (3, 2, "GooD GooD GooD GooD GooD GooD GooD GooD")`)
// 	if err != nil {
// 		fmt.Println("Failed to add rows to tables:", err)
// 		return nil, err
// 	}
