package model

import "time"

type Access struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Document
	AccessType
	ToWhom User
}
