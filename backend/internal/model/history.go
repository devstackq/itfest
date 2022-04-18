package model

import "time"

type History struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Document
	TypeEvent
	Author User
	ToWhom User
}
