package models

import "time"

type Tweet struct {
	ID        int
	Message   string
	CreatedAt time.Time
}

func NewTweet(message string) *Tweet {
	return &Tweet{
		Message: message,
	}
}
