package models

type Comment struct {
	Id            string `json:"id" validate:"required"`
	Creator       string `json:"creator" validate:"required"`
	PostId        string `json:"postId" validate:"required"`
	CommentInHtml string `json:"commentInHtml" validate:"required"`
	CreatedAt     string `json:"createdAt" validate:"required"`
}