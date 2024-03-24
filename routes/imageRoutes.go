package routes

import (
	"github.com/Abedmuh/tweeter-clone/middleware"

	"github.com/Abedmuh/tweeter-clone/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ImageRoutes(route *gin.RouterGroup, validate *validator.Validate) {
	ImageController := controllers.NewImageController()

	route.POST("/image", middleware.Authentication(),ImageController.PostImage)
}