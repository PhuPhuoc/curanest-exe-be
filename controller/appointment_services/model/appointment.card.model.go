package appointmentmodel

type AppointmentCardForCustomer struct {
	ID              string `db:"id" json:"id"`
	AppointmentDate string `db:"appointment_date" json:"appointment_date"`
	TimeFromTo      string `db:"time_from_to" json:"time_from_to"`
	Techniques      string `db:"techniques" json:"techniques"`
	TotalFee        int    `db:"total_fee" json:"total_fee"`
	Status          string `db:"status" json:"status"`
	NurseName       string `db:"full_name" json:"nurse_name"`
	NursePhone      string `db:"phone_number" json:"phone_number"`
	Avatar          string `db:"avatar" json:"avatar"`
}

type AppointmentCardForNurse struct {
	ID              string `db:"id" json:"id"`
	AppointmentDate string `db:"appointment_date" json:"appointment_date"`
	TimeFromTo      string `db:"time_from_to" json:"time_from_to"`
	Status          string `db:"status" json:"status"`
	Techniques      string `db:"techniques" json:"techniques"`
	TotalFee        int    `db:"total_fee" json:"total_fee"`
	PatientName     string `db:"full_name" json:"patient_name"`
	PatientPhone    string `db:"phone_number" json:"patient_phone"`
	Avatar          string `db:"avatar" json:"avatar"`
}
