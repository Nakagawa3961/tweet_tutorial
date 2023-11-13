package models

import "time"

type Tweet struct {
	ID        int
	UserID    int
	Text      string
	CreatedAt time.Time
}

func NewTweet(userID int, text string) *Tweet {
	return &Tweet{
		Text: text,
	}
}
