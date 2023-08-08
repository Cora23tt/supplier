package online_diler

type User struct {
	ID       int    `json:"-" db:"id"`
	Fullname string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}
