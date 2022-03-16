package service

import (
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
	DeleteUsers(id int, password string) error
}

type ProjectList interface {
}

type Item interface {
}

type Service struct {
	Authorization
	ProjectList
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
