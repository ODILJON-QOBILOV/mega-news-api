package models

import "time"

type Advertisement struct {
	Id        int
	Title     string
	Link      string
	Image     string
	CreatedAt time.Time
}