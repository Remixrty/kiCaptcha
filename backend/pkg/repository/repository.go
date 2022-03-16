package repository

import (
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
	DeleteUser(id int, password string) error
}

type ProjectList interface {
}

type Item interface {
}

type Repository struct {
	Authorization
	ProjectList
	Item
}

func NewRepositori(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
