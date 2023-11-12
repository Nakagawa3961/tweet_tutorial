package usecases

import (
	"github.com/ryuta06012/tweet_backend/src/models"
	"github.com/ryuta06012/tweet_backend/src/repositories"
)

type TweetInteractor struct {
	TweetRepository *repositories.TweetRepository
}

func NewTweetInteractor(tweetRepo *repositories.TweetRepository) *TweetInteractor {
	return &TweetInteractor{
		TweetRepository: tweetRepo,
	}
}

func (interactor *TweetInteractor) SaveTweet(tweet *models.Tweet) error {
	return interactor.TweetRepository.SaveTweet(tweet)
}
