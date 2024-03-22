package routes

import (
	"database/sql"
	"tweet-clone/controllers"
	"tweet-clone/middleware"
	"tweet-clone/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := service.NewPostService()
	controler := controllers.NewPostController(service, db,validate)

	path := route.Group("/post")
	path.Use(middleware.Authentication())
	{
		path.POST("/", controler.PostPost)
		path.GET("/", controler.GetPosts)
		
		path.POST("/comment", controler.PostComment)		
	}
}