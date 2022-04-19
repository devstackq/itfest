package model

import "time"

type User struct {
	ID          int
	Email       string
	Password    string
	FullName    string
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Company     `json:"company"`
	Departament `json:"departament"`
	Position    `json:"position"`
	Role        `json:"role"`
}

// todo:  Validation, Sanitaze
