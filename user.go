package online_diler

type User struct {
	ID        int    `json:"-" db:"id"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binging:"required"`
}
