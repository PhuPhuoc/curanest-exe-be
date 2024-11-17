package patientmodel

type PatientModelView struct {
	PatientModel
	Techniques []TechniqueModel `json:"techniques"`
}

type PatientModel struct {
	ID          string `db:"id" json:"id"`
	Avatar      string `db:"avatar" json:"avatar"`
	FullName    string `db:"full_name" json:"full_name"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Old         int    `db:"old" json:"old"`
	DOB         string `db:"dob" json:"dob"`
	CitizenID   string `db:"citizen_id" json:"citizen_id"`
	Ward        string `db:"ward" json:"ward"`
	District    string `db:"district" json:"district"`
	City        string `db:"city" json:"city"`
	Address     string `db:"address" json:"address"`
}

type TechniqueModel struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
