package paymentrepository

import (
	"fmt"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
)

func (store *paymentStore) GetAllWalletTransactionOfNurse(user_id ...string) ([]paymentmodel.WalletTransactionModel, error) {
	payments := []paymentmodel.WalletTransactionModel{}

	var query string
	if user_id != nil {
		query = `
		select id, user_id, amount, type, related_transaction_id, appointment_id, detail, created_at
		from wallet_transactions
		where user_id=?
		order by created_at desc
	`
		if err := store.db.Select(&payments, query, user_id); err != nil {
			return nil, fmt.Errorf("cannot get transactions <%w>", err)
		}
	} else {
		query = `
		select id, user_id, amount, type, related_transaction_id, appointment_id, detail, created_at
		from wallet_transactions
		order by created_at desc
	`
		if err := store.db.Select(&payments, query); err != nil {
			return nil, fmt.Errorf("cannot get transactions <%w>", err)
		}
	}
	return payments, nil
}
