package routes

import (
	"crud-auth-go/controllers"
	"crud-auth-go/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := service.NewUserService()
	controler := controllers.NewUserController(service, db,validate)

	path := route.Group("/user")
	{
		path.POST("/register", controler.PostUser)
		path.POST("/login", controler.PostLogin)
	}
}