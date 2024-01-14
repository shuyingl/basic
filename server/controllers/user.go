package controllers

import (
	"log"
	"net/http"
	"server/constants"
	"server/daos"
	"server/types"
	"server/utils"

	"server/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	LoginHandler(context *gin.Context)
	LogoutHandler(context *gin.Context)
	GetCurrentUserHandler(context *gin.Context)
}

type userController struct {
	baseController BaseController
	userDao        daos.UserDao
}

func NewUserController(userDao daos.UserDao) UserController {
	return &userController{&baseController{}, userDao}
}

func (uc *userController) LoginHandler(c *gin.Context) {
	socialMedia := types.SocialMedia(c.Param(constants.SOCIAL_MEDIA))
	if !socialMedia.IsValid() {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constants.INVALID_SOCIAL_MEDIA_PARAMETER))
		return
	}

	if socialMedia == types.Google {
		gToken := c.Request.Header.Get(constants.AUTHORIZATION)

		user, err := utils.GetUserFromGoogleToken(gToken, constants.GOOGLE_USERINFO_ENDPOINT)
		if err != nil || user.Email == "" {
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(constants.INVALID_GOOGLE_ACCESS_TOKEN))
			return
		}

		userT, err := uc.userDao.GetUser(&models.User{Email: user.Email})
		if err != nil {
			log.Printf("Creating new user: %v", user)
			user, err = uc.userDao.CreateUser(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error()))
				return
			}
		} else {
			user.ID = userT.ID
		}

		session := sessions.Default(c)
		session.Set(constants.AUTHENTICATED, true)
		session.Set(constants.USER_ID, user.ID.String())
		session.Save()

		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, models.User{})
	}
}

func (uc *userController) LogoutHandler(c *gin.Context) {
	_, err := uc.baseController.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error()))
		return
	}

	session := sessions.Default(c)

	// Revoke users authentication
	session.Set(constants.AUTHENTICATED, false)
	session.Save()

	c.Status(http.StatusOK)
}

func (uc *userController) GetCurrentUserHandler(c *gin.Context) {
	user, err := uc.baseController.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
