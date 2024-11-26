package paymentrepository

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// ProcessVNPayIPN processes the VNPay IPN request
func (store *paymentStore) ProcessVNPayIPN(vnpParams url.Values) error {
	// Configurations (Replace with your real configs)
	secretKey := "7BB0ONRZHMDTEIV7UKU0FT65SBLFESFK"

	// Lấy vnp_SecureHash từ query và xóa nó khỏi vnpParams
	secureHash := vnpParams.Get("vnp_SecureHash")
	vnpParams.Del("vnp_SecureHash")
	vnpParams.Del("vnp_SecureHashType") // Xóa nếu có

	// Sắp xếp các tham số theo thứ tự tăng dần
	keys := make([]string, 0, len(vnpParams))
	for key := range vnpParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Xây dựng raw query string đã sắp xếp
	var rawData strings.Builder
	for _, key := range keys {
		values := vnpParams[key]
		for _, value := range values {
			rawData.WriteString(key + "=" + value + "&")
		}
	}
	// Xóa dấu & cuối cùng
	signData := strings.TrimSuffix(rawData.String(), "&")

	// Tạo SecureHash
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(signData))
	expectedHash := hex.EncodeToString(h.Sum(nil))

	// So sánh hash
	if secureHash != expectedHash {
		return errors.New("checksum mismatch")
	}

	// Lấy thông tin đơn hàng từ tham số
	orderID := vnpParams.Get("vnp_TxnRef")
	rspCode := vnpParams.Get("vnp_ResponseCode")

	// Kiểm tra logic đơn hàng và cập nhật trạng thái
	// TODO: Thêm logic kiểm tra và cập nhật trạng thái đơn hàng dựa trên orderID và rspCode

	fmt.Println("orderID: ", orderID)
	fmt.Println("rspCode: ", rspCode)
	return nil
}
