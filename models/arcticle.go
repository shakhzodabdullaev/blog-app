package models

import "time"

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Article struct {
	ID        int        `json:"id"`
	Content              //Promoted Field
	AuthourID int        `json:"author_id"` //Nested struct
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ArticleListItem struct {
	ID        int        `json:"id"`
	Content              //Promoted Field
	Author    Person     `json:"author"` //Nested struct
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json: "updated_at"`
}

type ArticleCreatedModel struct {
	Content      //promoted Field
	AuthorID int `json:"author_id"`
}

type ArticleUpdateModel struct {
	ID int `json:"-"`
	Content
	AuthorID int `json:"author_id"`
}
