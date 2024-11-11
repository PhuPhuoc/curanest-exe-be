package patientmodel

type PatientCreationModel struct {
	ID          string   `db:"id" json:"-"`
	Avatar      string   `db:"avatar" json:"avatar"`
	FullName    string   `db:"full_name" json:"full_name"`
	Old         int      `db:"old" json:"-"`
	DOB         string   `db:"dob" json:"dob"`
	CitizenID   string   `db:"citizen_id" json:"citizen_id"`
	Address     string   `db:"address" json:"address"`
	PhoneNumber string   `db:"phone_number" json:"phone_number"`
	Techniques  []string ` json:"techniques"`
}
