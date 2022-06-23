package service

import (
	entity "DATABASECRUD/Entity"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserIface interface {
	Register(user *entity.User) (*entity.User, error)
	Login(user *entity.User, tempPassword string) (*entity.User, error)
	GetToken(id uint, email string, password string) string
	CheckToken(compareToken string, id uint, email string, password string) error
	UpdateUser(user *entity.User) (*entity.User, error)
	VerivyToken(tempToken string) string
}

type PhotoIface interface {
}

type UserSvc struct{}
type PhotoSvc struct{}

func NewPhotoSvc() PhotoIface {
	return &PhotoSvc{}
}

func NewUserSvc() UserIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User) (*entity.User, error) {
	// validasi field field user
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if len(user.Password) < 6 {
		return nil, errors.New("password must be minimum 6 characters")
	}
	if user.Age < 8 {
		return nil, errors.New("age must be greater than 8")
	}
	fmt.Println("udh dicek usernya")
	return user, nil
}

func (u *UserSvc) Login(user *entity.User, tempPassword string) (*entity.User, error) {
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	password := []byte(tempPassword)
	//check password salah
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), password); err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (u *UserSvc) UpdateUser(user *entity.User) (*entity.User, error) {
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	return user, nil
}

func (u *UserSvc) GetToken(id uint, email string, password string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(password))

	return signedToken
}

func (u *UserSvc) CheckToken(compareToken string, id uint, email string, password string) error {
	token := u.GetToken(id, email, password)
	if compareToken == token {
		fmt.Println("berhasil")
		return nil
	} else {
		fmt.Println("tidak berhasil")
		return errors.New("username cannot be empty")
	}
	//compare
}

func (u *UserSvc) VerivyToken(TempToken string) string {
	tokenString := TempToken
	claims := jwt.MapClaims{}
	var verivykey []byte
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return verivykey, nil
	})

	fmt.Println(token.Claims.Valid())
	return "nil"
}
