package models

import "time"

type Banner struct {
	Id          int
	Image       string
	Title       string
	Description string
	CreatedAt   time.Time
}

type BannerCarousel struct {
	Id          int
	Image       string
	Title       string
	Description string
	CreatedAt   time.Time
}