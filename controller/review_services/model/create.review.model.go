package reviewmodel

type CreateNewReviewModel struct {
	ID            string `db:"id" json:"-"`
	AppointmentID string `db:"appointment_id" json:"-" `
	PatientName   string `db:"patient_name" json:"patient_name" binding:"required"`
	Rate          int    `db:"rate" json:"rate" binding:"required,min=1,max=5"`
	Techniques    string `db:"techniques" json:"techniques"`
	Content       string `db:"content" json:"content"`
	CreatedAt     string `db:"created_at" json:"-"`
}
