package paymentmodel

type PaymentTransaction struct {
	OrderID       string  `db:"order_id" json:"order_id"`
	UserID        string  `db:"user_id" json:"user_id"`
	Amount        float64 `db:"amount" json:"amount"`
	OrderInfo     string  `db:"order_info" json:"order_info"`
	TransactionNo string  `db:"transaction_no" json:"transaction_no"`
	ResponseCode  string  `db:"response_code" json:"response_code"`
	PaymentStatus string  `db:"payment_status" json:"payment_status"`
	CreatedAt     string  `db:"created_at" json:"created_at"`
}
