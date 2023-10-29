package usecases

import (
	"github.com/ryuta06012/tweet_backend/src/models"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"github.com/ryuta06012/tweet_backend/src/views"
)

type UsersInteractor struct {
	UsersRepository *repositories.UsersRepository
}

func NewUsersInteractor(mysqlHandler *repositories.MysqlRepository) *UsersInteractor {
	return &UsersInteractor{
		UsersRepository: repositories.NewUsersRepository(mysqlHandler),
	}
}

func (u *UsersInteractor) CreateUsersRecord(request views.UserSignupRequest) error {
	user := models.NewUser(request.Name, request.Password)
	err := u.UsersRepository.CreateUserRecord(user)
	if err != nil {
		return err
	}
	return nil
}
