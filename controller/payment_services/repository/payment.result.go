package paymentrepository

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (store *paymentStore) VerifyVNPayResponse(vnpParams url.Values, secureHash, secretKey string) (bool, error) {
	// Create a new map to store the parameters
	params := make(map[string]string)

	// Convert url.Values to map[string]string, taking the first value for each key
	for key, values := range vnpParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	// Sort keys
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Build query string in the same way as GenerateVNPayURL
	var queryString strings.Builder
	for _, key := range keys {
		// Use the same URL encoding as in GenerateVNPayURL
		queryString.WriteString(fmt.Sprintf("%s=%s&", key, url.QueryEscape(params[key])))
	}

	// Remove trailing "&"
	signData := strings.TrimSuffix(queryString.String(), "&")

	// Calculate hash
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(signData))
	expectedHash := hex.EncodeToString(h.Sum(nil))

	// Compare hashes
	if secureHash != expectedHash {
		return false, fmt.Errorf("checksum mismatch: expected %s, got %s", expectedHash, secureHash)
	}

	return true, nil
}

// HandlePaymentResult processes the payment result based on orderID and response code
func (store *paymentStore) HandlePaymentResult(orderID, transactionNo, rspCode, amount string) (err error) {
	var tx *sqlx.Tx
	tx, err = store.db.Beginx()
	if err != nil {
		return fmt.Errorf("cannot begin transaction <%w>", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction <%w>", commitErr)
		}
	}()
	// TODO: Thêm logic kiểm tra và cập nhật trạng thái đơn hàng dựa trên orderID và rspCode
	query := `
	update transactions 
	set 
		transaction_no = ?,
		response_code = ?,
		payment_status = case 
			when ? = '00' then 'completed'
			else 'failed'
		end,
		last_updated = NOW()
	WHERE order_id = ?
`
	if _, err = tx.Exec(query, transactionNo, rspCode, rspCode, orderID); err != nil {
		return fmt.Errorf("cannot update status transactions <%w>", err)
	}

	// todo 1: create wallet_transaction
	// todo 1.1: get user_id from transactions
	query_select_user_id := `
		select user_id from transactions where order_id=?
	`
	var user_id string
	if err = tx.Get(&user_id, query_select_user_id, orderID); err != nil {
		return fmt.Errorf("get user_id from current transactions <%w>", err)
	}

	var floatAmount float64
	floatAmount, err = strconv.ParseFloat(amount, 64)
	floatAmount = floatAmount / 100
	if err != nil {
		return fmt.Errorf("invalid amount in transactions <%w>", err)
	}
	wallet_model := &paymentmodel.WalletTransaction{
		ID:                   uuid.New().String(),
		UserID:               user_id,
		Amount:               floatAmount,
		Type:                 "deposit",
		RelatedTransactionID: orderID,
		Detail:               "Tiền Nạp vào ví curanest",
		CreatedAt:            utils.GetCurrentDateTime(),
	}
	query_insert_wallet := `
		insert into wallet_transactions (
			id, user_id, amount, type, related_transaction_id, detail, created_at
		) values (
			:id, :user_id, :amount, :type, :related_transaction_id, :detail, :created_at
		)
	`
	if _, err = tx.NamedExec(query_insert_wallet, wallet_model); err != nil {
		return fmt.Errorf("cannot insert new wallet transaction <%w>", err)
	}
	// todo 2: update wallet_amount trong table users

	query_update_wallet_amount := `
		update users
		set
			wallet_amount = wallet_amount + ?
		where 
			id=?
	`
	if _, err = tx.Exec(query_update_wallet_amount, floatAmount, user_id); err != nil {
		return fmt.Errorf("cannot update amount in wallet <%w>", err)
	}
	return nil
}
