package model

import "time"

type Document struct {
	ID        int
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Author    User
}
