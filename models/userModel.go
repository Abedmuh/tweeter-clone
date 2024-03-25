package models

type Status string

const (
	Email Status = "email"
	Phone Status = "phone"
)

// main
type User struct {
	Id          string  `json:"id"`
	Name        string  `json:"name" validate:"required, min=5, max=50"`
	Password    string  `json:"password" validate:"required, min=5, max=15"`
	Phone       *string `json:"phone" validate:"min=5,max=15"`
	Email       *string `json:"email" validate:"min=5,max=15"`
	ImageUrl    *string `json:"imageUrl"`
	FriendCount *uint32 `json:"friendCount"`
}

// request
type UserRegister struct {
	CredentialsType   Status `json:"credentialType" validate:"required,oneof=email phone"`
	CredentialsValues string `json:"credentialValue" validate:"required"`
	Name              string `json:"name" validate:"required,min=5,max=50"`
	Password          string `json:"password" validate:"required,min=5,max=15"`
}

type UserLogin struct {
	CredentialsType   Status `json:"credentialType" validate:"required,oneof=email phone"`
	CredentialsValues string `json:"credentialValues" validate:"required,min=5,max=50"`
	Password          string `json:"password" validate:"required,min=5,max=50"`
}

type ReqUpEmail struct {
	Email string `json:"email" validate:"required"`
}

type ReqUpPhone struct {
	Phone string `json:"phone" validate:"required"`
}

type ReqPatchUser struct {
	Name     string `json:"name" validate:"required"`
	ImageUrl string `json:"imageUrl" validate:"required"`
}

// respon
type UserResLog struct {
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Name        string  `json:"name"`
	AccessToken string  `json:"accessToken"`
}

type ResRegUser struct {
	Id          string  `json:"id"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Name        string  `json:"name"`
	AccessToken string  `json:"accessToken"`
}
