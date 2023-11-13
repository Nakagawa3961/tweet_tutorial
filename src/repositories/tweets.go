package repositories

import (
	"github.com/ryuta06012/tweet_backend/src/models"
	"gorm.io/gorm"
)

type TweetRepository struct {
	Client *gorm.DB
}

func NewTweetRepository(mysqlHandler *MysqlRepository) *TweetRepository {
	return &TweetRepository{
		Client: mysqlHandler.Client,
	}
}

func (t *TweetRepository) CreateTweetRecord(tweet *models.Tweet) error {
	err := t.Client.Create(tweet).Error
	if err != nil {
		return err
	}
	return nil
}
