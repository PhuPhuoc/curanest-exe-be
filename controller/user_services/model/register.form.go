package usermodel

type RegisterForm struct {
	Name            string `db:"name" json:"name"`
	Email           string `db:"email" json:"email"`
	Password        string `db:"password" json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
