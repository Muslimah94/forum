package models

// PostDTO ...
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

// AuthorDTO ...
type AuthorDTO struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

// CommentDTO ...
type CommentDTO struct {
	ID         int    `json:"id"`
	AuthorID   int    `json:"author_id"`
	PostID     int    `json:"post_id"`
	Content    string `json:"content"`
	AuthorNick string `json:"author_nick"`
	Likes      int    `json:"likes"`
	Dislikes   int    `json:"dislikes"`
}

// CategoriesDTO ...
type CategoriesDTO struct {
	AllCategories []string `json:"all_categories"`
}
