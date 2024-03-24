package controllers

import (
	"database/sql"
	"tweet-clone/models"
	"tweet-clone/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostCtrlInter interface {
	PostPost(c *gin.Context)
	GetPosts(c *gin.Context)
	PostComment(c *gin.Context)
}

type PostController struct {
	PostService service.PostSvcInter
	DB 					*sql.DB
	validate 		*validator.Validate
}

func NewPostController(PostService service.PostSvcInter,DB *sql.DB, validate *validator.Validate) PostCtrlInter {
	return &PostController{
    PostService: PostService,
    DB: DB,
    validate: validate,
  }
}

func (p *PostController) PostPost(c *gin.Context) {
	var req models.ReqPost
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := p.validate.Struct(req); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	if err := p.PostService.AddPost(req, c, p.DB); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
    "message": "successfully add post",
  })
}

func (p *PostController) GetPosts(c *gin.Context) {
	posts, err := p.PostService.GetAllPost(c, p.DB)
	if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
		return
  }
	c.JSON(200, gin.H{
    "posts": posts,
  })
}

func (p *PostController) PostComment(c *gin.Context) {
	var req models.ReqComment
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := p.validate.Struct(req); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	if err := p.PostService.CheckPost(req.PostId, c, p.DB); err!= nil {
    c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
  }

  if err := p.PostService.AddComment(req, c, p.DB); err!= nil {
    c.JSON(400, gin.H{
			"message": "fail to add comment",
			"error": err.Error(),
		})
		return
  }
  c.JSON(200, gin.H{
    "message": "successfully add comment",
  })
}
