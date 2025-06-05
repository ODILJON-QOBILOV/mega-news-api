package models

import "time"

type Comment struct {
	Comment string `json:"comment"`
	User            User
	News            News
	Created_at      time.Time
	Updated_at      time.Time
}
