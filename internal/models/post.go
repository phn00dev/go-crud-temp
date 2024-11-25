package models

import "time"

type Post struct {
	Id        int
	Title     string
	Content   string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
