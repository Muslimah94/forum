package models

import "encoding/json"

// PostDTO ...
type PostDTO struct {
	ID           int             `json:"id"`
	Author       AuthorDTO       `json:"author"`
	Title        string          `json:"title"`
	Content      json.RawMessage `json:"content"`
	Categories   []string        `json:"categories"`
	Likes        int             `json:"likes"`
	Dislikes     int             `json:"dislikes"`
	Comments     int             `json:"comments"`
	CategoriesID []int           `json:"categories_id"`
	CreationDate int64           `json:"creation_date"`
}

// AuthorDTO ...
type AuthorDTO struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

// CommentDTO ...
type CommentDTO struct {
	ID       int       `json:"id"`
	Author   AuthorDTO `json:"author"`
	PostID   int       `json:"post_id"`
	Content  string    `json:"content"`
	Likes    int       `json:"likes"`
	Dislikes int       `json:"dislikes"`
}

// CategoriesDTO ...
type CategoriesDTO struct {
	AllCategories   []string `json:"all_categories"`
	NewCategoryName string   `json:"new_category_name"`
}

type ReactionDTO struct {
	ID        int `json:"id"`
	AuthorID  int `json:"author_id"`
	Type      int `json:"type"`
	PostID    int `json:"post_id"`
	CommentID int `json:"comment_id"`
}

type RegisterUser struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Error struct {
	Status      string `json:"status"`
	Description string `json:"desciption"`
}
