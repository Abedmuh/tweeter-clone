package models

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