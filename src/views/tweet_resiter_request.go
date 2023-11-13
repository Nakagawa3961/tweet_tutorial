package views

type TweetRequest struct {
	UserID int    `json:"userId" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
