package service

import (
	online_diler "github.com/cora23tt/online-diler"
	"github.com/cora23tt/online-diler/pkg/repository"
)

type Authorisation interface {
	CreateUser(user online_diler.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorisation
	EAuthService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		EAuthService:  NewEAuthService(),
	}
}
