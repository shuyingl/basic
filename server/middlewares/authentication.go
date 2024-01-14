package middlewares

import (
	"net/http"
	"server/constants"
	"server/daos"
	"server/utils"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var URLS_NO_NEED_LOGIN = []string{
	"/api/v1/user/login/social",
}

func AuthMiddleware(
	userDao daos.UserDao,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		fullPath := c.FullPath()
		var noNeedLogin = false
		for _, url := range URLS_NO_NEED_LOGIN {
			if strings.HasPrefix(fullPath, url) {
				noNeedLogin = true
				break
			}
		}

		if !noNeedLogin {
			session := sessions.Default(c)

			// Check if user is authenticated
			if auth, ok := session.Get(constants.AUTHENTICATED).(bool); !ok || !auth {
				c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constants.INVALID_TOKEN))
				c.Abort()
				return
			}

			user, err := userDao.GetUserWithId(session.Get(constants.USER_ID).(string))
			if err != nil {
				c.JSON(http.StatusUnauthorized, utils.ErrorMessage(err.Error()))
				c.Abort()
				return
			}

			c.Set(constants.CURRENT_USER, user)
		}

		c.Next()
	}
}
