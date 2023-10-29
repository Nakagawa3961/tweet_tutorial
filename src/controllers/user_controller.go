package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryuta06012/tweet_backend/src/repositories"
	"github.com/ryuta06012/tweet_backend/src/usecases"
	"github.com/ryuta06012/tweet_backend/src/views"
)

type UsersController struct {
	interactor *usecases.UsersInteractor
}

func NewUsersController(mysqlHandler *repositories.MysqlRepository) *UsersController {
	return &UsersController{
		interactor: usecases.NewUsersInteractor(mysqlHandler),
	}
}

func (u *UsersController) RegisterFrontUser(c *gin.Context) {
	var request views.UserSignupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		views.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	if err := u.interactor.CreateUsersRecord(request); err != nil {
		views.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
