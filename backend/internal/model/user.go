package model

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	FullName  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Company
	Departament
	Position
	Role
}

// todo:  Validation, Sanitaze
