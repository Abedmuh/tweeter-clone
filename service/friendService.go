package service

import (
	"database/sql"
	"errors"

	"github.com/Abedmuh/tweeter-clone/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FriendSvcInter interface {
	AddFriend(friend models.ReqFriend, c *gin.Context, tx *sql.DB) (models.Friend, error)
	DeleteFriend(friend models.ReqFriend,c *gin.Context, tx *sql.DB)  error
	GetFriends (c *gin.Context, tx *sql.DB) ([]models.ResFriend, error)
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
	if newFriend.UserId == newFriend.FriendId {
		return models.Friend{}, errors.New("cant add yourself as friend")
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
  result, err := tx.Exec(query, userId, friend.UserId)
	if err != nil {
		return err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("friend not found")
	}
  return nil
}

func (f *FriendService) GetFriends(c *gin.Context, tx *sql.DB) ([]models.ResFriend, error) {
	user,_ := c.Get("user")
	userId := string(user.(string)) 
	var friends []models.ResFriend

  query:= `SELECT users.id, users.name, users.imageUrl, users.friendCount, friends.created_at
		FROM friends
		JOIN users ON friends.friend_id = users.id
		WHERE friends.user_id = $1
		ORDER BY friends.created_at
	`
  rows, err := tx.Query(query, userId)
  if err!= nil {
    return friends, err
  }
  for rows.Next() {
    var user models.ResFriend
    err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.ImageUrl,
		  &user.FriendCount,
			&user.CreatedAt)
    if err!= nil {
      return friends, err
    }
    friends = append(friends, user)
  }
  return friends, nil
}