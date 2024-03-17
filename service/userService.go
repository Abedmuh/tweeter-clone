package service

import (
	"crud-auth-go/models"
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserSvcInter interface {
	AddUser(user models.UserRegister, c *gin.Context, tx *sql.DB) (models.User, error)
	Login(uuserLogin models.UserLogin, userDb models.User, c *gin.Context, tx *sql.DB) (string, error)
	CheckUser(user string, c *gin.Context, tx *sql.DB)  (models.User, error)
}

type UserService struct {
}

func NewUserService() UserSvcInter{
	return &UserService{}
}

func (us *UserService) AddUser(user models.UserRegister, c *gin.Context, tx *sql.DB) (models.User, error) {
	var newUser models.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err!= nil {
    return newUser, err
  }

	query:= `INSERT INTO users (name, password, phone, email)
		VALUES ($1, $2, $3, $4) 
		RETURNING id, name, password, phone, email
	`
	err = tx.QueryRow(query,
		newUser.Name, 
		hashedPassword, 
		newUser.Phone, 
		newUser.Email).Scan(
			&newUser.Id, 
			&newUser.Name, 
			&newUser.Password, 
			&newUser.Phone, 
			&newUser.Email)
	
	if err!= nil {
    return newUser, err
  }

  return newUser, nil
}

func (us *UserService) Login(userLogin models.UserLogin, userDb models.User , c *gin.Context, tx *sql.DB) (string, error) {
	
	err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(userDb.Password))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "not match error"})
		return "",errors.New("password salah")
	}

	secretKey := viper.GetString("JWT_SECRET_KEY")
	timeExp := viper.GetDuration("JWT_TIME_EXP")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userDb.Id,
    "exp": time.Now().Add(time.Duration(timeExp) * time.Hour).Unix(), 
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

  return signedToken, nil
}

func (us *UserService) CheckUser(user string, c *gin.Context, tx *sql.DB) (models.User, error) {
	var newUser models.User


	query := `SELECT id, name, password, phone, email FROM users WHERE phone = $1`
	err := tx.QueryRow(query, user).Scan(
    &newUser.Id, 
    &newUser.Name, 
    &newUser.Password, 
    &newUser.Phone,
		&newUser.Email)
	
  if err!= nil {
    return models.User{}, err
  }
  return newUser, nil
}