package appointmentmodel

type DetailAppointment struct {
	AppointmentModel   `json:"appointment_information"`
	NurseInformation   NurseInfo    `json:"nurse_information"`
	PatientInformation PatientModel `json:"patient_infomation"`
}

type AppointmentModel struct {
	ID              string `db:"id" json:"-"`
	NurseID         string `db:"nurse_id" json:"-"`
	PatientID       string `db:"patient_id" json:"-"`
	AppointmentDate string `db:"appointment_date" json:"appointment_date"`
	TimeFromTo      string `db:"time_from_to" json:"time_from_to"`
	Status          string `db:"status" json:"status"`
	Techniques      string `db:"techniques" json:"techniques"`
	TotalFee        int    `db:"total_fee" json:"total_fee"`
}

type NurseInfo struct {
	NurseName   string `db:"full_name" json:"nurse_name"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Avatar      string `db:"avatar" json:"avatar"`
}

type PatientModel struct {
	Avatar             string `db:"avatar" json:"avatar"`
	FullName           string `db:"full_name" json:"full_name"`
	PhoneNumber        string `db:"phone_number" json:"phone_number"`
	Old                int    `db:"old" json:"old"`
	DOB                string `db:"dob" json:"dob"`
	Ward               string `db:"ward" json:"ward"`
	District           string `db:"district" json:"district"`
	City               string `db:"city" json:"city"`
	Address            string `db:"address" json:"address"`
	MedicalDescription string `db:"medical_description" json:"medical_description"`
	NoteForNurses      string `db:"note_for_nurses" json:"note_for_nurses"`
}
