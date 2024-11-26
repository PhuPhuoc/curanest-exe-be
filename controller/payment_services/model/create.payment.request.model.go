package paymentmodel

type CreatePaymentRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}
