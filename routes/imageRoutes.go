package routes

import (
	"crud-auth-go/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ImageRoutes(route *gin.RouterGroup, validate *validator.Validate) {
	ImageController := controllers.NewImageController()

	route.POST("/image", ImageController.PostImage)
}