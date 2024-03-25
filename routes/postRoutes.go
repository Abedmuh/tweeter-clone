package routes

import (
	"database/sql"

	"github.com/Abedmuh/tweeter-clone/middleware"

	"github.com/Abedmuh/tweeter-clone/controllers"
	"github.com/Abedmuh/tweeter-clone/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := service.NewPostService()
	controler := controllers.NewPostController(service, db,validate)

	path := route.Group("/post")
	{
		path.POST("/",middleware.Authentication(), controler.PostPost)
		path.GET("/", controler.GetPosts)
		path.GET("/:id",controler.GetUserPosts)
		path.PUT("/:id",middleware.Authentication(), controler.PutPost)
		path.DELETE("/:id",middleware.Authentication(), controler.DeletePost)

		
		path.GET("/user/:id", controler.GetPost)
		path.POST("/:id/comment",middleware.Authentication(), controler.PostComment)
		path.DELETE("/comment/:id",middleware.Authentication(), controler.DeleteComment)
	}
}