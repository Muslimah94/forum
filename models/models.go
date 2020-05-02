package models

// Users ...
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

// Roles ...
type Roles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Categories ...
type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoriesDTO struct {
	AllCategories []string `json:"all_categories"`
}

// Posts ...
type Post struct {
	ID           int
	AuthorID     int
	Title        string
	Content      string
	CreationDate int64
}

type PostDTO struct {
	ID           int       `json:"id"`
	Author       AuthorDTO `json:"author"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Categories   []string  `json:"categories"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	Comments     int       `json:"comments"`
	CategoriesID []int     `json:"categories_id"`
	CreationDate int64     `json:"creation_date"`
}

type AuthorDTO struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

// Comments ...
type Comments struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
}

type CommentsDTO struct {
	AuthorNick string `json:"author_nick"`
	Likes      int    `json:"likes"`
	Dislikes   int    `json:"dislikes"`
}

// Reactions ...
type Reactions struct {
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
type PostCategoriesDTO struct {
}

//
