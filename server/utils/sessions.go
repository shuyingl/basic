package utils

import (
	"context"
	"encoding/json"
	"io"
	"server/constants"
	"server/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
)

func GetSessionHandler() gin.HandlerFunc {
	address, password, secretKey, maxAge := "localhost:6379", "", "secret-key", 3600

	// In a local environment, no need to specify the domain. It will default to the origin of the request.
	domainOption := ""

	store, err := redis.NewStore(10, "tcp", address, password, []byte(secretKey))
	if err != nil {
		panic(err.Error())
	}

	store.Options(sessions.Options{
		Domain:   domainOption,
		Path:     "/",
		MaxAge:   maxAge,
		Secure:   false, // Should be true in production when using HTTPS
		HttpOnly: true,
	})

	return sessions.Sessions(constants.TOKEN, store)
}

func GetUserFromGoogleToken(token string, googleUserInfoEndpoint string) (*models.User, error) {
	// create a http.client from the access token
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	// make a request to the userinfo endpoint with the client
	resp, err := client.Get(googleUserInfoEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the http response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user *models.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return user, nil
}
