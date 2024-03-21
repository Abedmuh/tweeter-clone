package main

import (
	"crud-auth-go/config"
	"crud-auth-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, err := config.GetDBConnection()
	if err!= nil {
    panic(err)
  }
	defer db.Close()

	gin := gin.Default()
	validate := validator.New()

	v1 := gin.Group("v1")
	{
		routes.UserRoutes(v1,db,validate)
		routes.FriendRoutes(v1,db,validate)
	}

	gin.Run(":8000")

}