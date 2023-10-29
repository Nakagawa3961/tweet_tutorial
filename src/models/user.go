package models

import "time"

type User struct {
	ID        int
	Name      string
	Password string
	CreatedAt time.Time
}

func NewUser(name string, password string) *User {
	return &User{
		Name: name,
		Password: password,
	}
}
