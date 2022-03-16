package service

import (
	"crypto/sha1"

	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/repository"
)

//parametes in token
const (
	salt      = "klfjalsghwehhih2315r8342yfbq321" //future in .env
	tokenTTL  = 12 * time.Hour
	signinkey = "lfailgifsdkhafqp3284jr21h4u09"
)

//struct for jwt token
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

//struct for method for users,
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

//create user with hash passord
//-----------------------------
func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

// func (s *AuthService) DeleteUser(user model.User) (int, error) {
// 	user.Password = s.generateHashPassword(user.Password)
// 	return s.repo.DeleteUser(user)
// }

//Create Token using jwt
//-------------
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generateHashPassword(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			// IssuedAt:  time.Now().Unix(),
		})
	// int(user.ID),

	return token.SignedString([]byte(signinkey))
}

//Create hash password
//--------------------
func (s *AuthService) generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

//Delete user with hash password
//-----------------------------
func (s *AuthService) DeleteUsers(id int, password string) error {
	Hashpassword := s.generateHashPassword(password)
	err := s.repo.DeleteUser(id, Hashpassword)
	return err
}
