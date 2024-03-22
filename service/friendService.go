package service

import (
	"database/sql"
	"tweet-clone/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FriendSvcInter interface {
	AddFriend(friend models.ReqFriend, c *gin.Context, tx *sql.DB) (models.Friend, error)
	DeleteFriend(friend models.ReqFriend,c *gin.Context, tx *sql.DB)  error
	GetFriends (c *gin.Context, tx *sql.DB) ([]models.Friend, error)
}

type FriendService struct {
}

func NewFriendService() FriendSvcInter{
  return &FriendService{}
}

func (f *FriendService) AddFriend(friend models.ReqFriend, c *gin.Context, tx *sql.DB) (models.Friend, error) {
	
	user, _ := c.Get("user")
	id := uuid.New().String()

	newFriend := models.Friend{
		Id: id,
		UserId: string(user.(string)),
    FriendId: friend.UserId,
	}

  query:= `INSERT INTO friends (id,user_id, friend_id)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, friend_id
  `
  err := tx.QueryRow(query,
		newFriend.Id,
		newFriend.UserId, 
		newFriend.FriendId).Scan(
			&newFriend.Id, 
      &newFriend.UserId, 
      &newFriend.FriendId,
		)
  if err!= nil {
    return newFriend, err
  }
  return newFriend, nil
}

func (f *FriendService) DeleteFriend(friend models.ReqFriend,c *gin.Context, tx *sql.DB)  error {

	user,_ := c.Get("user")
	userId := string(user.(string)) 

  query:= `DELETE FROM friends WHERE user_id = $1 AND friend_id = $2`
  _, err := tx.Exec(query, userId, friend.UserId)
  if err!= nil {
    return err
  }
  return nil
}

func (f *FriendService) GetFriends(c *gin.Context, tx *sql.DB) ([]models.Friend, error) {
	user,_ := c.Get("user")
	userId := string(user.(string)) 
	var friends []models.Friend

  query:= `SELECT * FROM friends WHERE user_id = $1`
  rows, err := tx.Query(query, userId)
  if err!= nil {
    return friends, err
  }
  for rows.Next() {
    var friend models.Friend
    err := rows.Scan(
			&friend.Id, 
			&friend.UserId, 
			&friend.FriendId)
    if err!= nil {
      return friends, err
    }
    friends = append(friends, friend)
  }
  return friends, nil
}