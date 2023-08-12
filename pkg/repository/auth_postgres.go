package repository

import (
	online_diler "github.com/cora23tt/online-diler"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	DB *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{DB: db}
}

func (s *AuthPostgres) CreateUser(user online_diler.User) (int, error) {
	// 	var id int

	// 	parts := strings.Split(user.Fullname, " ")
	// 	if len(parts) != 2 {
	// 		return 0, errors.New("INVALID FULLNAME")
	// 	}
	// 	firstName := parts[1]
	// 	lastName := parts[0]

	// 	query := `INSERT INTO users
	// 	(username, password_hash, first_name, last_name, email, address, telephone)
	// 	VALUES($1, $2, $3, $4, $5, '', 0) RETURNING id`

	// 	row := s.DB.QueryRow(query, user.Username, user.Password, firstName, lastName, user.Email)
	// 	if err := row.Scan(&id); err != nil {
	// 		return 0, err
	// 	}
	id := 0
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (online_diler.User, error) {
	// var user online_diler.User
	// query := "SELECT id FROM users WHERE username=$1 AND password_hash=$2"
	// err := r.DB.Get(&user, query, username, password)
	user := online_diler.User{}
	var err error
	return user, err
}
