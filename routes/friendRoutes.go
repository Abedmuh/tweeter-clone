package routes

import (
	"crud-auth-go/controllers"
	"crud-auth-go/middleware"
	"crud-auth-go/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FriendRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := service.NewFriendService()
	controler := controllers.NewFriendController(service, db,validate)

	path := route.Group("/friend")
	path.Use(middleware.Authentication())
	{
		path.POST("/", controler.PostFriend)
		path.GET("/", controler.GetFriends)
		path.DELETE("/", controler.DeleteFriend)
	}
}