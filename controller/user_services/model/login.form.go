package usermodel

type LoginForm struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
