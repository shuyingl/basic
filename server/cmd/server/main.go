package main

import (
	"flag"
	"fmt"
	"log"
	"server/controllers"
	"server/daos"
	"server/middlewares"
	"server/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Importing the driver for PostgreSQL
	"gorm.io/gorm"
)

func main() {
	portPtr := flag.Int("port", 8000, "port to run the server on")
	dbNamePtr := flag.String("db_name", "test", "name of the database to create")
	flag.Parse()

	router, db := setupRouter(*dbNamePtr)
	defer utils.ClosePostgresConnection(db)

	err := router.Run(fmt.Sprintf(":%d", *portPtr))
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func setupRouter(dbName string) (
	*gin.Engine,
	*gorm.DB,
) {
	var (
		db *gorm.DB = utils.GetPostgresConnection(
			dbName,
		)
		userDao        daos.UserDao               = daos.NewUserDao(db)
		userController controllers.UserController = controllers.NewUserController(
			userDao,
		)
		sessionHandler gin.HandlerFunc = utils.GetSessionHandler()
		authMiddleware gin.HandlerFunc = middlewares.AuthMiddleware(
			userDao,
		)
	)

	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	apiV1.Use(sessionHandler)
	apiV1.Use(middlewares.LoggingMiddleware())

	apiV1.Use(authMiddleware)
	{
		user := apiV1.Group("/user")
		{
			user.GET("/login/social/:social_media", userController.LoginHandler)
			user.GET("/logout", userController.LogoutHandler)
			user.GET("/current", userController.GetCurrentUserHandler)
		}
	}
	return router, db
}
