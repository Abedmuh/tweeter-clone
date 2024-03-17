package controllers

import (
	"crud-auth-go/models"
	"crud-auth-go/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlInter interface {
	PostUser(c *gin.Context)
	PostLogin(c *gin.Context)
}

type UserController struct {
	UserService service.UserSvcInter
	DB 					*sql.DB
	validate 		*validator.Validate
}

func NewUserController(userService service.UserSvcInter,DB *sql.DB, validate *validator.Validate) UserCtrlInter {
	return &UserController{
    UserService: userService,
    DB: DB,
    validate: validate,
  }
}

func (u *UserController) PostUser(c *gin.Context) {
	var user models.UserRegister
  if err := c.ShouldBindJSON(&user); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  if err := u.validate.Struct(user); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  if _,err := u.UserService.CheckUser(user.CredentialsValues,c, u.DB); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	//main
	newUser, err := u.UserService.AddUser(user,c,u.DB)
	if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  c.JSON(200, gin.H{
		"message": "User Created",
	  "data": newUser,
  })
}

func (u *UserController) PostLogin(c *gin.Context) {
	var user models.UserLogin
  if err := c.ShouldBindJSON(&user); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  if err := u.validate.Struct(user); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	
	newUser, err := u.UserService.CheckUser(string(user.CredentialsValues),c, u.DB)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	//main
  token, err := u.UserService.Login(user,newUser,c,u.DB)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	result := models.UserResLog{
		Email: newUser.Email,
		Phone: newUser.Phone,
		Name: newUser.Name,
		AccessToken: token,
	}

	c.JSON(200, gin.H{
    "message": "User Created",
    "data": result,
  })
}