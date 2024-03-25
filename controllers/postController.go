package controllers

import (
	"database/sql"

	"github.com/Abedmuh/tweeter-clone/models"
	"github.com/Abedmuh/tweeter-clone/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostCtrlInter interface {
	PostPost(c *gin.Context)
	GetPosts(c *gin.Context)
	GetUserPosts(c *gin.Context)
	PutPost(c *gin.Context)
	DeletePost(c *gin.Context)
	
	GetPost(c *gin.Context)
	PostComment(c *gin.Context)
	DeleteComment(c *gin.Context)
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

func (p *PostController) GetUserPosts(c *gin.Context) {
	posts, err := p.PostService.GetUserPosts(c, p.DB)
	if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
		return
  }
	c.JSON(200, gin.H{
    "posts": posts,
  })
}
func (p *PostController) GetPosts(c *gin.Context) {
	posts, err := p.PostService.GetPosts(c, p.DB)
	if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
		return
  }
	c.JSON(200, gin.H{
    "posts": posts,
  })
}

func (p *PostController) PutPost(c *gin.Context) {
  if err := p.PostService.AuthoPost(c, p.DB); err!= nil {
    c.AbortWithStatusJSON(404, gin.H{
      "error": err.Error(),
    })
    return
  }

  var req models.ReqPost
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := p.validate.Struct(req); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
	res,err := p.PostService.UpdatePost(req, c, p.DB)
	if err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
	c.JSON(200, gin.H{
    "message": "successfully update post",
		"post": res,
  })
}

func (p *PostController) DeletePost(c *gin.Context) {
  if err := p.PostService.AuthoPost(c, p.DB); err!= nil {
    c.AbortWithStatusJSON(404, gin.H{
      "error": err.Error(),
    })
    return
  }

  if err := p.PostService.DeletePost(c, p.DB); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
    "message": "successfully delete post",
  })
}

func (p *PostController) GetPost(c *gin.Context) {
  post, err := p.PostService.GetPost(c, p.DB)
	if err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  comments, err := p.PostService.GetComments(post.Id, c, p.DB)
	if err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }

	c.JSON(200, gin.H{
    "post": post,
		"comments": comments,
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

func (p *PostController) DeleteComment(c *gin.Context) {
	id := c.Param("id")

  if err := p.PostService.AuthoComment(id, c, p.DB); err!= nil {
    c.AbortWithStatusJSON(404, gin.H{
      "error": err.Error(),
    })
    return
  }

  if err := p.PostService.DeleteComment(id, c, p.DB); err!= nil {
    c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
    "message": "successfully delete comment",
  })
}