package appointmentmodel

type WalletTransactionModel struct {
	ID            string  `db:"id" json:"id"`
	UserID        string  `db:"user_id" json:"user_id"`
	Amount        float64 `db:"amount" json:"amount"`
	Type          string  `db:"type" json:"type"`
	AppointmentID string  `db:"appointment_id" json:"appointment_id"`
	Detail        string  `db:"detail" json:"detail"`
	CreatedAt     string  `db:"created_at" json:"created_at"`
}
