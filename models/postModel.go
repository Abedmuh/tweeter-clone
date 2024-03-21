package models

//inDatabase
type Post struct {
	Id         string    `json:"id" validate:"required"`
	Creator    string    `json:"creator" validate:"required"`
	PostInHtml string    `json:"postInHtml" validate:"required"`
	Tags       []string  `json:"tags"`
	CommentId  []string  `json:"commentId"`
	CreatedAt  string    `json:"createdAt" validate:"required"`
	UpdatedAt  string    `json:"updatedAt" validate:"required"`
}