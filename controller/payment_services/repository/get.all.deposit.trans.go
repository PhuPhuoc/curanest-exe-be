package paymentrepository

import (
	"fmt"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
)

func (store *paymentStore) GetAllDepositTransactionOfNurse(user_id ...string) ([]paymentmodel.PaymentTransactionModel, error) {
	payments := []paymentmodel.PaymentTransactionModel{}

	var query string
	if user_id != nil {
		query = `
		select order_id, user_id, amount, order_info, transaction_no, response_code, payment_status, created_at
		from transactions
		where user_id=?
		order by created_at desc
	`
		if err := store.db.Select(&payments, query, user_id); err != nil {
			return nil, fmt.Errorf("cannot get transactions <%w>", err)
		}
	} else {
		query = `
		select order_id, user_id, amount, order_info, transaction_no, response_code, payment_status, created_at
		from transactions
		order by created_at desc
	`
		if err := store.db.Select(&payments, query); err != nil {
			return nil, fmt.Errorf("cannot get transactions <%w>", err)
		}
	}

	return payments, nil
}
