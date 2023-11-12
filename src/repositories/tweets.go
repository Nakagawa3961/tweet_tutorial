package repositories

import (
	"github.com/ryuta06012/tweet_backend/src/models"
	"gorm.io/gorm"
)

type TweetRepository struct {
	Client *gorm.DB
}

func NewTweetRepository(db *gorm.DB) *TweetRepository {
	return &TweetRepository{
		Client: db,
	}
}

func (repo *TweetRepository) SaveTweet(tweet *models.Tweet) error {

	err := repo.Client.Create(tweet).Error
	if err != nil {

		return err
	}

	return nil
}
