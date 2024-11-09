package customermodel

type CustomerModel struct {
	UserID      string `db:"user_id" json:"-"`
	CitizenID   string `db:"citizen_id" json:"citizen_id"`
	FullName    string `db:"full_name" json:"full_name"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	DOB         string `db:"dob" json:"dob"`
	Ward        string `db:"ward" json:"ward"`
	District    string `db:"district" json:"district"`
	City        string `db:"city" json:"city"`
	Address     string `db:"address" json:"address"`
}
