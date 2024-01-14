package controllers

import (
	"errors"
	"server/constants"

	"server/models"

	"github.com/gin-gonic/gin"
)

type BaseController interface {
	GetCurrentUser(c *gin.Context) (*models.User, error)
}

type baseController struct{}

func (bc *baseController) GetCurrentUser(c *gin.Context) (*models.User, error) {
	var user, ok = c.Get(constants.CURRENT_USER)
	if !ok {
		return nil, errors.New(constants.FAILED_TO_GET_CURRENT_USER)
	}

	userT, ok := user.(*models.User)
	if !ok {
		return nil, errors.New(constants.FAILED_TO_GET_CURRENT_USER)
	}
	return userT, nil
}
