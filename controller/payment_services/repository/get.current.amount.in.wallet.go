package paymentrepository

import "fmt"

type WalletAmountModel struct {
	WalletAmount float64 `db:"wallet_amount" json:"wallet_amount"`
}

func (store *paymentStore) GetWalletAmount(user_id string) (*WalletAmountModel, error) {
	wallet := &WalletAmountModel{}
	query := `
		select wallet_amount from users where id=?
	`
	if err := store.db.Get(wallet, query, user_id); err != nil {
		return nil, fmt.Errorf("cannot get amount in user wallet <%w>", err)
	}
	return wallet, nil
}
