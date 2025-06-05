package models

import "time"

type Follow struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	WriterId    int `json:"writer_id"`
	CreatedAt   time.Time
    UpdatedAt   time.Time
}