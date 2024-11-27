package paymentrepository

import (
	"fmt"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
)

func (store *paymentStore) GetAllDepositTransactionOfNurse(user_id string) ([]paymentmodel.PaymentTransaction, error) {
	payments := []paymentmodel.PaymentTransaction{}

	query := `
		select order_id, user_id, amount, order_info, transaction_no, response_code, payment_status, created_at
		from transactions
		where user_id=?
	`
	if err := store.db.Select(&payments, query, user_id); err != nil {
		return nil, fmt.Errorf("cannot get transactions <%w>", err)
	}
	return payments, nil
}
