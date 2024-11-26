package usermodel

type User struct {
	ID           string  `db:"id" json:"id"`
	Email        string  `db:"email" json:"email"`
	Password     string  `db:"password" json:"-"`
	Name         string  `db:"name" json:"user_name"`
	Avatar       string  `db:"avatar" json:"avatar"`
	Role         string  `db:"role" json:"role"`
	WalletAmount float64 `db:"wallet_amount" json:"-"`
	CreatedAt    string  `db:"created_at" json:"-"`
}
