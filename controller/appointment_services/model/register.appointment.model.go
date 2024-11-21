package appointmentmodel

type RegisterAppoinment struct {
	AppointmentModel
	ListNurseWorkSchedules []string
}

type AppointmentModel struct {
	ID              string `db:"id" json:"-"`
	PatientID       string `db:"patient_id" json:"patient_id"`
	NurseID         string `db:"nurse_id" json:"nurse_id"`
	AppointmentDate string `db:"appointment_date" json:"appointment_date"`
	TimeFromTo      string `db:"time_from_to" json:"time_from_to"`
	Status          string `db:"status" json:"-"`
	Techniques      string `db:"techniques" json:"techniques"`
	TotalFee        int    `db:"total_fee" json:"total_fee"`
}
