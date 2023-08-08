package service

import (
	online_diler "github.com/cora23tt/online-diler"
	"github.com/cora23tt/online-diler/pkg/repository"
)

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user online_diler.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	return "", nil
}

func (*AuthService) ParseToken(accessToken string) (int, error) {
	UserID := 1
	return UserID, nil
}

func generatePasswordHash(password string) string {
	return ""
}
