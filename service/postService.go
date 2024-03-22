package service

import (
	"crud-auth-go/models"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostSvcInter interface {
	AddPost(req models.ReqPost,c *gin.Context, tx *sql.DB) error
	GetAllPost(c *gin.Context, tx *sql.DB) ([]models.Post, error)

	AddComment(req models.ReqComment,c *gin.Context, tx *sql.DB) error
}

type PostService struct {}

func NewPostService() PostSvcInter{
  return &PostService{}
}

func (ps *PostService) AddPost(req models.ReqPost,c *gin.Context, tx *sql.DB) error {

	id := uuid.New().String()
	user,_ := c.Get("user")
	creator := string(user.(string)) 
	
  query := `INSERT INTO posts (id, creator, post_in_html, tags) VALUES ($1, $2, $3, $4)`
  _, err := tx.QueryContext(c, query, 
		id,
		creator, 
		req.PostInHtml,
	  pq.Array(req.Tags))
  if err!= nil {
    return err
  }
  return nil
}

func (ps *PostService) GetAllPost(c *gin.Context, tx *sql.DB) ([]models.Post, error) {
	var posts []models.Post

	user,_ := c.Get("user")
	creator := string(user.(string))

	query := `SELECT * FROM posts WHERE creator = $1`

	row, err := tx.QueryContext(c, query, creator)
	if err!= nil {
    return posts, err
  }
	defer row.Close()
	// slide window technice
	for row.Next() {
		var post models.Post
    err := row.Scan(
			&post.Id, 
			&post.Creator, 
			&post.PostInHtml, 
			pq.Array(&post.Tags),
			&post.CommentId,
      &post.CreatedAt,
		)
    if err!= nil {
      return posts, err
    }
    posts = append(posts, post)
	}
	return posts, nil
}

func (ps *PostService) AddComment(req models.ReqComment,c *gin.Context, tx *sql.DB) error {
	id := uuid.New().String()
  user,_ := c.Get("user")
  creator := string(user.(string)) 

  query := `INSERT INTO comments (id, creator, post_id, comment_in_html) VALUES ($1, $2, $3, $4)`
  _, err := tx.QueryContext(c, query, 
    id,
    creator,
    req.PostId,
    req.Comment)
  if err!= nil {
    return err
  }
  return nil
}
