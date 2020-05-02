package models

// Post ...
type Post struct {
	ID           int
	AuthorID     int
	Title        string
	Content      string
	CreationDate int64
}

// User ...
type User struct {
	ID       int
	Email    string
	Nickname string
	Password string
	RoleID   int
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
	AuthorID  int
	Type      int
	PostID    int
	CommentID int
}

// PostCategories ...
type PostCategories struct {
	PostID       int
	CategoryID   int
	CategoryName string
}
