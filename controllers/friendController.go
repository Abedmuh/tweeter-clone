package controllers

import (
	"database/sql"
	"tweet-clone/models"
	"tweet-clone/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FriendCtrlInter interface {
	PostFriend(c *gin.Context)
	GetFriends(c *gin.Context)
	DeleteFriend(c *gin.Context)
}

type FriendController struct {
	FriendService service.FriendSvcInter
	DB 					*sql.DB
	validate 		*validator.Validate
}

func NewFriendController(FriendService service.FriendSvcInter,DB *sql.DB, validate *validator.Validate) FriendCtrlInter {
	return &FriendController{
    FriendService: FriendService,
    DB: DB,
    validate: validate,
  }
}

func (f *FriendController) PostFriend(c *gin.Context) {
	var friend models.ReqFriend
  if err := c.ShouldBindJSON(&friend); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
	}
  if err := f.validate.Struct(friend); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	
  newFriend, err := f.FriendService.AddFriend(friend, c, f.DB)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
		"message": "berhasil",
	  "data": newFriend,
  })
}

func (f *FriendController) GetFriends(c *gin.Context) {
  friend, err := f.FriendService.GetFriends(c, f.DB)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(400, gin.H{
		"message": "berhasil mendapatkan teman",
		"data": friend,
  })
}

func (f *FriendController) DeleteFriend(c *gin.Context) {
	var friend models.ReqFriend
	if err := c.ShouldBindJSON(&friend); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
	}

	if err := f.validate.Struct(friend); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  err := f.FriendService.DeleteFriend(friend, c, f.DB)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	c.JSON(400, gin.H{
		"message": "berhasil menghapus teman",
  })

}
