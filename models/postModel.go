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
	PostId string `json:"postId" validate:"required"`
	Comment string `json:"comment" validate:"required,min=2,max=500"`
}

type subUser struct {
	Id          string  `json:"userId"`
	Name        string  `json:"name" validate:"required, min=5, max=50"`
	ImageUrl    *string `json:"imageUrl"`
	FriendCount *uint32 `json:"friendCount"`
}

type subPost struct {
	PostInHtml string `json:"postInHtml" validate:"required,min=2,max=500"`
	Tags       []string `json:"tags" validate:"required"`
	CreatedAt  string `json:"createdAt" validate:"required"`
}

type subComments struct {
	Comment 		string `json:"comment" validate:"required"`
	Creator     subUser `json:"creator" validate:"required"`
	CreatedAt   string `json:"createdAt" validate:"required"`
}
//respons
type ResPost struct {
	PostId  string `json:"postId" validate:"required"`
	Post 		subPost `json:"post"`
	Comment subComments `json:"comment" validate:"required,min=2,max=500"`
	Creator ReqFriend	`json:"creator" validate:"required"`
}