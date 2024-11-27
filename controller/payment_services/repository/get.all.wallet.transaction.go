package paymentrepository

import (
	"fmt"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
)

func (store *paymentStore) GetAllWalletTransactionOfNurse(user_id string) ([]paymentmodel.WalletTransactionModel, error) {
	payments := []paymentmodel.WalletTransactionModel{}

	query := `
		select id, user_id, amount, type, related_transaction_id, appointment_id, detail, created_at
		from wallet_transactions
		where user_id=?
	`
	if err := store.db.Select(&payments, query, user_id); err != nil {
		return nil, fmt.Errorf("cannot get transactions <%w>", err)
	}
	return payments, nil
}
