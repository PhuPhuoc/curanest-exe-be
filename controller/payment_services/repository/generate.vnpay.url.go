package paymentrepository

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
)

func (store *paymentStore) GenerateVNPayURL(req *paymentmodel.CreatePaymentRequest, vnp_config *paymentmodel.VnpayConfig) (string, error) {
	orderID := fmt.Sprintf("%s_%s", req.UserID, time.Now().Format("150405"))
	createDate := time.Now().Format("20060102150405")
	orderInfo := "Nạp tiền vào ví curanest"

	initialPayment := paymentmodel.PaymentTransaction{
		OrderID:       orderID,
		UserID:        req.UserID,
		Amount:        float64(req.Amount),
		OrderInfo:     orderInfo,
		PaymentStatus: "pending",
		CreatedAt:     utils.GetCurrentDateTime(),
	}

	if err := store.CreateInitialPayment(initialPayment); err != nil {
		return "", fmt.Errorf("failed to create initial payment record: %v", err)
	}

	// Xử lý số tiền (nhân 100)
	amount := req.Amount * 100

	vnpParams := map[string]string{
		"vnp_Version":    "2.1.0",
		"vnp_Command":    "pay",
		"vnp_TmnCode":    vnp_config.TMNCode,
		"vnp_Amount":     fmt.Sprintf("%d", amount),
		"vnp_CurrCode":   "VND",
		"vnp_TxnRef":     orderID,
		"vnp_OrderInfo":  orderInfo,
		"vnp_OrderType":  "billpayment",
		"vnp_Locale":     "vn",
		"vnp_CreateDate": createDate,
		"vnp_ReturnUrl":  vnp_config.ReturnURL,
		"vnp_IpAddr":     "127.0.0.1",
	}

	// Sắp xếp các tham số theo thứ tự tăng dần
	keys := make([]string, 0, len(vnpParams))
	for key := range vnpParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Xây dựng query string theo thứ tự
	var queryString strings.Builder
	for _, key := range keys {
		queryString.WriteString(fmt.Sprintf("%s=%s&", key, url.QueryEscape(vnpParams[key])))
	}
	rawQuery := strings.TrimSuffix(queryString.String(), "&")

	// Tạo SecureHash
	h := hmac.New(sha512.New, []byte(vnp_config.SecretKey))
	h.Write([]byte(rawQuery))
	secureHash := hex.EncodeToString(h.Sum(nil))
	vnpParams["vnp_SecureHash"] = secureHash

	// Build final URL
	query := url.Values{}
	for k, v := range vnpParams {
		query.Add(k, v)
	}
	finalURL := fmt.Sprintf("%s?%s", vnp_config.VNPURL, query.Encode())
	return finalURL, nil
}

func (store *paymentStore) CreateInitialPayment(payment paymentmodel.PaymentTransaction) error {
	query := `
        insert into transactions (
            order_id, user_id, amount, order_info, payment_status, created_at
        ) values (
            :order_id, :user_id, :amount, :order_info, :payment_status, :created_at
        )
    `
	_, err := store.db.NamedExec(query, payment)
	return err
}
