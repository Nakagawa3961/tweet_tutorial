package repositories

import (
	"github.com/ryuta06012/tweet_backend/src/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	Client *gorm.DB
}

func NewUsersRepository(mysqlHandler *MysqlRepository) *UsersRepository {
	return &UsersRepository{
		Client: mysqlHandler.Client,
	}
}

func (u *UsersRepository) CreateUserRecord(user *models.User) error {
	err := u.Client.Create(user).Error

	if err != nil {
		return err
	}
	return nil
}
