package models

type Category struct {
	Id int
	Title string `json:"title"`
	Image string `json:"image"`
}