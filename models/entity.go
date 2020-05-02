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
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

// Role ...
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Category ...
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comments ...
type Comment struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
}

// Reactions ...
type Reaction struct {
	ID        int `json:"id"`
	AuthorID  int `json:"author_id"`
	Type      int `json:"type"`
	PostID    int `json:"post_id"`
	CommentID int `json:"comment_id"`
}

// PostsCategories ...
type PostCategories struct {
	PostID       int
	CategoryID   int
	CategoryName string
}
