package models

import "time"

type Tweet struct {
	ID        int
	UserID    int
	Text      string
	CreatedAt time.Time
}

func NewTweet(userId int, text string) *Tweet {
	return &Tweet{
		UserID: userId,
		Text:   text,
	}
}
