package service

import (
	"database/sql"
	"errors"

	"github.com/Abedmuh/tweeter-clone/models"
	"github.com/Abedmuh/tweeter-clone/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserSvcInter interface {
	AddUser(user models.UserRegister, c *gin.Context, tx *sql.DB) (models.ResRegUser, error)
	RegistCheck(user string, c *gin.Context, tx *sql.DB) error
	
	LoginUserCheck(user string, c *gin.Context, tx *sql.DB) (models.User,error)
	Login(user models.UserLogin,userdb models.User , c *gin.Context, tx *sql.DB) (models.UserResLog, error)
	
	PatchEmail(req models.ReqUpEmail,c *gin.Context, tx *sql.DB) error
	PatchPhone(req models.ReqUpPhone,c *gin.Context, tx *sql.DB) error

	PatchUser(req models.ReqPatchUser,c *gin.Context, tx *sql.DB) error
}

type UserService struct {
}

func NewUserService() UserSvcInter{
	return &UserService{}
}

func (us *UserService) AddUser(user models.UserRegister, c *gin.Context, tx *sql.DB) (models.ResRegUser, error) {
	var newUser models.User
	var resUser models.ResRegUser
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err!= nil {
    return resUser, err
  }
	userid := uuid.New().String()

	var query string
	if user.CredentialsType == "email" {
		query = `INSERT INTO users (id, name, password, email)
			VALUES ($1, $2, $3, $4) 
			RETURNING id, name, password, phone, email
		`
	} else {
		query = `INSERT INTO users (id, name, password, phone)
			VALUES ($1, $2, $3, $4) 
			RETURNING id, name, password, phone, email
		`
	}

	err = tx.QueryRow(query,
		userid,
		user.Name, 
		hashedPassword, 
		user.CredentialsValues, 
	  ).Scan(
			&newUser.Id, 
			&newUser.Name, 
			&newUser.Password, 
			&newUser.Phone, 
			&newUser.Email)
	
	if err!= nil {
    return resUser, err
  }

	token,_:= utils.GenerateToken(userid)

	resUser = models.ResRegUser{
		Id: newUser.Id,
		Email: newUser.Email,
    Phone: newUser.Phone,
		Name: newUser.Name,
		AccessToken: token,
	}

  return resUser, nil
}

func (us *UserService) RegistCheck(user string, c *gin.Context, tx *sql.DB) error {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 OR phone = $1)`
	err := tx.QueryRow(query, user).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return errors.New("user already exists")
}

func (us *UserService) Login(userLogin models.UserLogin, userDb models.User , c *gin.Context, tx *sql.DB) (models.UserResLog, error) {

	err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password))
	if err != nil {
		return models.UserResLog{},errors.New("password salah")
	}

	token, err := utils.GenerateToken(userDb.Id)
	if err!= nil {
    return models.UserResLog{}, err
  }

	resLog := models.UserResLog{
		Email: userDb.Email,
    Phone: userDb.Phone,
    Name: userDb.Name,
    AccessToken: token,
	}

  return resLog, nil
}

func (us *UserService) LoginUserCheck(user string, c *gin.Context, tx *sql.DB) (models.User,error) {
	var userDb models.User
  query := `SELECT * FROM users WHERE email = $1 OR phone = $1`
  err := tx.QueryRow(query, user).Scan(
    &userDb.Id, 
    &userDb.Name, 
    &userDb.Password, 
    &userDb.Email,
    &userDb.Phone, 
	  &userDb.ImageUrl,
		&userDb.FriendCount)
  if err!= nil {
    return userDb, err
  }
  return userDb, nil
}

func (ps *UserService) PatchEmail(req models.ReqUpEmail,c *gin.Context, tx *sql.DB) error {
	user,_ := c.Get("user")
  creator := string(user.(string)) 

  query := `UPDATE users SET email = $1 WHERE id = $2`
  _, err := tx.ExecContext(c, query, req.Email, creator)
  if err!= nil {
    return err
  }
  return nil
}

func (ps *UserService) PatchPhone(req models.ReqUpPhone,c *gin.Context, tx *sql.DB) error {
	user,_ := c.Get("user")
  creator := string(user.(string)) 

  query := `UPDATE users SET phone = $1 WHERE id = $2`
  _, err := tx.ExecContext(c, query, req.Phone, creator)
  if err!= nil {
    return err
  }
  return nil
}

func (ps *UserService) PatchUser(req models.ReqPatchUser,c *gin.Context, tx *sql.DB) error {
	user,_ := c.Get("user")
  creator := string(user.(string)) 

  query := `UPDATE users SET name = $1, imageUrl =$2 WHERE id = $3`
  _, err := tx.ExecContext(c, query, 
		req.Name,
		req.ImageUrl, 
		creator)
  if err!= nil {
    return err
  }
  return nil
}


