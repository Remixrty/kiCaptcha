package repository

import (
	"errors"

	"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//Creating user with validation
//-----------------------------
func (p *AuthPostgres) CreateUser(user model.User) (int, error) {
	//check error when adding users
	var count int64
	if p.db.Model(&model.User{}).Where("email = ?", user.Email).Count(&count); count != 0 {
		err := errors.New("email is always used")
		return 0, err
	}
	if p.db.Model(&model.User{}).Where("phone = ?", user.Phone).Count(&count); count != 0 {
		err := errors.New("phone is always used")
		return 0, err
	}
	p.db.Create(&user)
	return int(user.Id), nil
}

//Getting users using username and passowrd with validation
//-----------------------------------------
func (p *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	g := p.db.Model(&model.User{}).First(&user, "email = ?", username)
	if user.Password != password {
		g.Error = errors.New("wrong password")
		return user, g.Error
	}

	return user, g.Error
}

//DeleteUser by id and password with validation
//-----------------------------
func (p *AuthPostgres) DeleteUser(id int, password string) error {
	//delete user in db
	var user model.User
	g := p.db.Model(&model.User{}).First(&user, "id = ?", id)

	if user.Password != password {
		g.Error = errors.New("wrong password")
		return g.Error
	}

	if g.Error == nil {
		p.db.Delete(&model.User{}, id)
	}

	return g.Error
}
