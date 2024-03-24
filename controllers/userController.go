package controllers

import (
	"database/sql"

	"github.com/Abedmuh/tweeter-clone/models"
	"github.com/Abedmuh/tweeter-clone/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlInter interface {
	PostUser(c *gin.Context)
	PostLogin(c *gin.Context)

	PostEmail(c *gin.Context)
	PostPhone(c *gin.Context)

	PatchUser(c *gin.Context)
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
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

  if err := u.validate.Struct(user); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
  }

	if user.CredentialsType == "email" {
		if err := u.validate.Var(user.CredentialsValues, "required,email"); err!= nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			panic(err)
		}
	} else {
		if err := u.validate.Var(user.CredentialsValues, "required"); err!= nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			panic(err)
		}
	}


  if err := u.UserService.RegistCheck(user.CredentialsValues,c, u.DB); err!= nil {
    c.AbortWithStatusJSON(409, gin.H{"error": err.Error()})
    return
  }

	//main
	newUser, err := u.UserService.AddUser(user,c,u.DB)
	if err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

  c.JSON(201, gin.H{
		"message": "User registered successfully",
	  "data": newUser,
  })
}

func (u *UserController) PostLogin(c *gin.Context) {
	var user models.UserLogin
  if err := c.ShouldBindJSON(&user); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

  if err := u.validate.Struct(user); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
	
	newUser, err := u.UserService.LoginUserCheck(string(user.CredentialsValues),c, u.DB)
  if err!= nil {
    c.AbortWithStatusJSON(404, gin.H{
			"message": "login user check",
			"error": err.Error(),
		})
    return
  }

	//main
  result, err := u.UserService.Login(user,newUser,c,u.DB)
  if err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	c.JSON(200, gin.H{
    "message": "User logged successfully",
    "data": result,
  })
}


func (u *UserController) PostEmail(c *gin.Context) {
	var req models.ReqUpEmail
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := u.validate.Struct(req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := u.UserService.PatchEmail(req, c, u.DB); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{
      "message": "fail to add email",
      "error": err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "successfully add email",
  })
}

func (u *UserController) PostPhone(c *gin.Context) {
	var req models.ReqUpPhone
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := u.validate.Struct(req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := u.UserService.PatchPhone(req, c, u.DB); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{
      "message": "fail to add phone",
      "error": err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "successfully add phone",
  })
}

func (u *UserController) PatchUser(c *gin.Context) {
	var req models.ReqPatchUser
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := u.validate.Struct(req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := u.UserService.PatchUser(req, c, u.DB); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{
      "message": "fail to patch user",
      "error": err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "successfully patch user",
  })
}
