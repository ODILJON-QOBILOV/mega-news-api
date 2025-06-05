package models

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	SecondImage string `json:"secondImage"`
	AuthorID    int    `json:"author_id"`
	Author      User   `json:"author" gorm:"foreignKey:AuthorID"`
}