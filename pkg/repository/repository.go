package repository

import (
	online_diler "github.com/cora23tt/online-diler"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user online_diler.User) (int, error)
	GetUser(username, password string) (online_diler.User, error)
}

type Repository struct {
	Authorisation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
