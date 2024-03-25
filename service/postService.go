package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Abedmuh/tweeter-clone/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostSvcInter interface {
	AddPost(req models.ReqPost,c *gin.Context, tx *sql.DB) error
	GetPosts(c *gin.Context, tx *sql.DB) ([]models.Post, error)
	GetUserPosts(c *gin.Context, tx *sql.DB) ([]models.Post, error)
	UpdatePost(req models.ReqPost , c *gin.Context, tx *sql.DB) (models.Post, error)
	DeletePost(c *gin.Context, tx *sql.DB) error
	
	GetPost(c *gin.Context, tx *sql.DB) (models.Post, error)
	GetComments(id string, c *gin.Context, tx *sql.DB) ([]models.Comment, error)
	AddComment(req models.ReqComment,c *gin.Context, tx *sql.DB) error
	DeleteComment(req string, c *gin.Context, tx *sql.DB) error

	AuthoComment(req string,c *gin.Context, tx *sql.DB) error
	AuthoPost(c *gin.Context, tx *sql.DB) error
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

func (ps *PostService) GetUserPosts(c *gin.Context, tx *sql.DB) ([]models.Post, error) {
	var posts []models.Post
	
	id := c.Param("id")
	query := `SELECT * FROM posts WHERE creator = $1`

	row, err := tx.QueryContext(c, query, id)
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
      &post.CreatedAt,
		)
    if err!= nil {
      return posts, err
    }
    posts = append(posts, post)
	}
	return posts, nil
}
func (ps *PostService) GetPosts(c *gin.Context, tx *sql.DB) ([]models.Post, error) {
	var posts []models.Post
	
	query := `SELECT * FROM posts`

	row, err := tx.QueryContext(c, query)
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
      &post.CreatedAt,
		)
    if err!= nil {
      return posts, err
    }
    posts = append(posts, post)
	}
	return posts, nil
}

func (ps *PostService) AuthoPost(c *gin.Context, tx *sql.DB) error {
	user,_ := c.Get("user")
	reqUser, ok := user.(string)
	if !ok {
		c.AbortWithStatusJSON(403, "Unathorized user")
		return errors.New("Unathorized")
	}
	param := c.Param("id")

	var owner string 
	query := `SELECT creator FROM posts WHERE id = $1`
	err := tx.QueryRowContext(c, query, param).Scan(&owner)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(404, "Not Found")
    	return errors.New("Unathorized")
		}
		return err
	}
	if owner != reqUser {
    c.AbortWithStatusJSON(403, "Unathorized user")
    return errors.New("Unathorized")
  }
	return nil
}

func (ps *PostService) UpdatePost(req models.ReqPost, c *gin.Context, tx *sql.DB) (models.Post, error) {
	var post models.Post
	param := c.Param("id")
  query := `UPDATE posts 
		SET post_in_html = $1, tags = $2 
		WHERE id = $3
		RETURNING *
	`

  err := tx.QueryRowContext(c, query,
    req.PostInHtml,
    pq.Array(req.Tags),
    param).Scan(
			&post.Id,
      &post.Creator,
      &post.PostInHtml,
      pq.Array(&post.Tags),
      &post.CreatedAt,
		)
  if err!= nil {
    return post, err
  }
  return post, nil
}

func (ps *PostService) DeletePost(c *gin.Context, tx *sql.DB) error {
	id := c.Param(":id")
	query := `DELETE FROM posts WHERE id = $1`
  _, err := tx.ExecContext(c, query, id)
  if err!= nil {
    return err
  }
  return nil
}
func (ps *PostService) AddComment(req models.ReqComment,c *gin.Context, tx *sql.DB) error {
	id := uuid.New().String()
  user,_ := c.Get("user")
  creator := string(user.(string))
	param := c.Param("id")

  query := `INSERT INTO comments (id, creator, post_id, comment_in_html) VALUES ($1, $2, $3, $4)`
  _, err := tx.QueryContext(c, query, 
    id,
    creator,
    param,
    req.Comment)
  if err!= nil {
    return err
  }
  return nil
}
func (ps *PostService) GetPost(c *gin.Context, tx *sql.DB) (models.Post, error) {
	var post models.Post
	param := c.Param("id")
  query := `SELECT * FROM posts WHERE id = $1 LIMIT 1`
	row := tx.QueryRowContext(c, query, param)
	err := row.Scan(
		&post.Id,
		&post.Creator,
		&post.PostInHtml,
		pq.Array(&post.Tags),
		&post.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
				return post, fmt.Errorf("post not found")
		}
		return post, err
	}
  return post, nil
}

func (ps *PostService) GetComments(id string, c *gin.Context, tx *sql.DB) ([]models.Comment, error) {
	var comments []models.Comment
  
  query := `SELECT * FROM comments WHERE post_id = $1`

  row, err := tx.QueryContext(c, query, id)
  if err!= nil {
    return comments, err
  }
  defer row.Close()
  // slide window technice
  for row.Next() {
    var comment models.Comment
    err := row.Scan(
      &comment.Id, 
      &comment.Creator, 
      &comment.PostId, 
      &comment.CommentInHtml,
      &comment.CreatedAt,
    )
    if err!= nil {
      return comments, err
    }
    comments = append(comments, comment)
  }
  return comments, nil
}

func (ps *PostService) AuthoComment(req string,c *gin.Context, tx *sql.DB) error {
	user,_ := c.Get("user")
	reqUser, ok := user.(string)
	if !ok {
		c.AbortWithStatusJSON(403, "Unathorized user")
		return errors.New("Unathorized")
	}

	var creator string
  query := `SELECT creator FROM comments WHERE id = $1`
  err := tx.QueryRowContext(c, query, req).Scan(&creator)
  if err != nil {
    return err
  }
	if creator != reqUser {
		c.AbortWithStatusJSON(403, "Unathorized user")
		return errors.New("Unathorized")
	}
  return nil
}

func (ps *PostService) DeleteComment(req string, c *gin.Context, tx *sql.DB) error {
	query := `DELETE FROM comments WHERE id = $1`
  _, err := tx.ExecContext(c, query, req)
  if err!= nil {
    return err
  }
  return nil
}