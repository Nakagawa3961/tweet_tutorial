package views

type TweetRequest struct {
	Message string `json:"message" binding:"required"`
}
