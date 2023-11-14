package usecases

import (
	"fmt"

	"github.com/ryuta06012/tweet_backend/src/models"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"github.com/ryuta06012/tweet_backend/src/views"
)

type TweetInteractor struct {
	TweetRepository *repositories.TweetRepository
}

func NewTweetInteractor(mysqlHandler *repositories.MysqlRepository) *TweetInteractor {
	return &TweetInteractor{
		TweetRepository: repositories.NewTweetRepository(mysqlHandler),
	}
}

func (t *TweetInteractor) CreateTweetsRecord(request views.TweetRequest) error {
	tweet := models.NewTweet(request.UserID, request.Text)
	fmt.Printf("tweet.UserID: %v\n", tweet.UserID)
	err := t.TweetRepository.CreateTweetRecord(tweet)
	if err != nil {
		return err
	}
	return nil
}
