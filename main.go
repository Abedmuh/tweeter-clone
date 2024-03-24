package main

import (
	"github.com/Abedmuh/tweeter-clone/config"
	"github.com/Abedmuh/tweeter-clone/middleware"
	"github.com/Abedmuh/tweeter-clone/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, err := config.GetDBConnection()
	if err!= nil {
    panic(err)
  }
	defer db.Close()

	app := gin.Default()
	app.Use(middleware.RecoveryMiddleware())
	validate := validator.New()

	v1 := app.Group("v1")
	{
		routes.UserRoutes(v1,db,validate)
		routes.FriendRoutes(v1,db,validate)
		routes.PostRoutes(v1,db,validate)
		routes.ImageRoutes(v1,validate)
	}
	app.Run(":8000")
}