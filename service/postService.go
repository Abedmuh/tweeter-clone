package service

import (
	"crud-auth-go/models"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type PostSvcInter interface {
	AddPost(post models.Post, c *gin.Context, tx *sql.DB) (models.Post, error)
	GetAllPost(id int, c *gin.Context, tx *sql.DB) ([]models.Post, error)
	UpdatePost(post models.Post, c *gin.Context, tx *sql.DB) (models.Post, error)
	DeletePost(id int, c *gin.Context, tx *sql.DB) (models.Post, error)
}

type PostService struct {}

func NewPostService() PostSvcInter{
  return &PostService{}
}

func (ps *PostService) AddPost(post models.Post, c *gin.Context, tx *sql.DB) (models.Post, error) {
	var newPost models.Post
  query := `INSERT INTO posts (id, creator, postInHtml, tags) VALUES ($1, $2, $3, $4)`
  _, err := tx.QueryContext(c, query, 
		post.Id,
		post.Creator, 
		post.PostInHtml,
	  pq.Array(post.Tags))
  if err!= nil {
    return newPost, err
  }
  return newPost, nil
}

func (ps *PostService) GetAllPost(id int, c *gin.Context, tx *sql.DB) ([]models.Post, error) {
	var posts []models.Post

	query := `SELECT * FROM post WHERE id = $1`

	row, err := tx.QueryContext(c, query, id)
	if err!= nil {
    return posts, err
  }
	defer row.Close()
	for row.Next() {
		var post models.Post
    err := row.Scan(
			&post.Id, 
			&post.Creator, 
			&post.PostInHtml, 
			pq.Array(&post.Tags),
			&post.CommentId,
      &post.CreatedAt,
			&post.UpdatedAt,
		)
    if err!= nil {
      return posts, err
    }
    posts = append(posts, post)
	}
	return posts, nil
}

func (ps *PostService) UpdatePost(post models.Post, c *gin.Context, tx *sql.DB) (models.Post, error) {
	var newPost models.Post
  query := `UPDATE posts SET postInHtml = $2, tags = $3 WHERE id = $1`
  _, err := tx.QueryContext(c, query, 
    post.Id,
    post.PostInHtml,
    pq.Array(post.Tags))
  if err!= nil {
    return newPost, err
  }
  return newPost, nil
}

func (ps *PostService) DeletePost(id int, c *gin.Context, tx *sql.DB) (models.Post, error) {
	var newPost models.Post
  query := `DELETE FROM posts WHERE id = $1`
  _, err := tx.QueryContext(c, query, id)
  if err!= nil {
    return newPost, err
  }
  return newPost, nil
}