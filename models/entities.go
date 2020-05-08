package models

import uuid "github.com/satori/go.uuid"

// Post ...
type Post struct {
	ID           int
	AuthorID     int
	Title        string
	Content      []byte
	CreationDate int64
}

// User ...
type User struct {
	ID        int
	Nickname  string
	FirstName string
	LastName  string
	Avatar    string
	RoleID    int
}

// Role ...
type Role struct {
	ID   int
	Name string
}

// Category ...
type Category struct {
	ID   int
	Name string
}

// Comment ...
type Comment struct {
	ID       int
	AuthorID int
	PostID   int
	Content  string
}

// Reaction ...
type Reaction struct {
	ID        int
	Type      int
	AuthorID  int
	PostID    int
	CommentID int
}

// PostCat ...
type PostCat struct {
	ID           int
	PostID       int
	CategoryID   int
	CategoryName string
}

// Session ...
type Session struct {
	ID      int
	UserID  int
	UUID    uuid.UUID
	ExpDate int64
}

// Credentials ...
type Credentials struct {
	ID             int
	Email          string
	HashedPassword string
}
