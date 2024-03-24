package models

import "time"

type (
	SortBy string
	OrderBy string
) 

const (
	SortByFriendCount SortBy = "friendCount"
	SortByCreatedAt   SortBy = "createdAt"
	OrderByAsc OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)


// main db
type Friend struct {
	Id       string `json:"id" validate:"required"`
	UserId   string `json:"name"`
	FriendId string `json:"friend"`
}

// request
type ReqFriend struct {
	UserId string `json:"userId" validate:"required"`
}

type ResFriend struct {
	Id          string  `json:"userId"`
	Name        string  `json:"name" validate:"required, min=5, max=50"`
	ImageUrl    *string `json:"imageUrl"`
	FriendCount *uint32 `json:"friendCount"`
	CreatedAt 	*time.Time `json:"createdAt"`
}

type ParamFriend struct {
	Limit      int    
	Offset     int    
	SortBy     SortBy 
	OrderBy    OrderBy
	OnlyFriend bool   
	Search     string 
}