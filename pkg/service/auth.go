package service

import (
	"encoding/base64"

	online_diler "github.com/cora23tt/online-diler"
	"github.com/cora23tt/online-diler/pkg/repository"
	"golang.org/x/crypto/bcrypt"
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

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}

func (*AuthService) ParseToken(accessToken string) (int, error) {
	UserID := 1
	return UserID, nil
}

func generatePasswordHash(password string) string {
	return ""
}
