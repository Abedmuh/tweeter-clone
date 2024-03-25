package models

//inDatabase
type Post struct {
	Id         string    `json:"id" validate:"required"`
	Creator    string    `json:"creator" validate:"required"`
	PostInHtml string    `json:"postInHtml" validate:"required"`
	Tags       []string  `json:"tags"`
	CreatedAt  string    `json:"createdAt" validate:"required"`

}

//request
type ReqPost struct {
	PostInHtml string `json:"postInHtml" validate:"required,min=2,max=500"`
	Tags       []string `json:"tags" validate:"required"`
}

type ReqComment struct {
	Comment string `json:"comment" validate:"required,min=2,max=500"`
}