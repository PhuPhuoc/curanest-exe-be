package paymenthandler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	paymentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

)

// https://curanest.com.vn/payment-result-success?amount=100000&date=25%2F11%2F2024&infor=Chuyen+khoan
// https://curanest.com.vn/payment-result-fail?amount=100000&date=25%2F11%2F2024&infor=Chuyen+khoan&response-code=02

func VNPayReturnHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			log.Println("(dev) Error loading .env file")
			redirectURL := fmt.Sprintf(
				"%s?amount=%s&date=%s&infor=%s&response-code=%s",
				"https://curanest.com.vn/payment-result-fail?",
				"0",
				"Hệ thống thanh toán hiện đang bảo trì - Vui lòng thử lại sau",
				"---",
				"---",
			)
			c.Redirect(http.StatusFound, redirectURL)
			return
		}
		secretKey := os.Getenv("VNPAY_SECRET_KEY")
		client_url_success := os.Getenv("CLIENT_URL_PAYMENT_SUCCESS")
		client_url_fail := os.Getenv("CLIENT_URL_PAYMENT_FAIL")

		payment_content := "Chuyển tiền vào ví Curanest"
		payment_date := utils.GetCurrentDateTime()

		// Lấy tất cả tham số từ query
		vnpParams := c.Request.URL.Query()

		// Lấy vnp_SecureHash và loại bỏ nó khỏi tham số
		secureHash := vnpParams.Get("vnp_SecureHash")
		vnpParams.Del("vnp_SecureHash")
		vnpParams.Del("vnp_SecureHashType")

		orderID := vnpParams.Get("vnp_TxnRef")
		rspCode := vnpParams.Get("vnp_ResponseCode")
		transactionNo := vnpParams.Get("vnp_TransactionNo")
		amount := vnpParams.Get("vnp_Amount")

		var floatAmount float64
		floatAmount, _ = strconv.ParseFloat(amount, 64)
		floatAmount = floatAmount / 100
		// Gọi repository để xác minh SecureHash
		repo := paymentrepository.NewPaymentStore(db)
		valid, err := repo.VerifyVNPayResponse(vnpParams, secureHash, secretKey)
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{"code": "97", "message": "Verification error: " + err.Error()})
			redirectURL := fmt.Sprintf(
				"%s?amount=%f&date=%s&infor=%s&response-code=%s",
				client_url_fail,
				floatAmount,
				payment_content+" - xác thực không thành công",
				payment_date,
				rspCode,
			)
			c.Redirect(http.StatusFound, redirectURL)
			return
		}

		// Nếu SecureHash không hợp lệ
		if !valid {
			// c.JSON(http.StatusOK, gin.H{"code": "97", "message": "Invalid checksum"})
			redirectURL := fmt.Sprintf(
				"%s?amount=%f&date=%s&infor=%s&response-code=%s",
				client_url_fail,
				floatAmount,
				payment_content+" - kiểm tra không hợp lệ",
				payment_date,
				rspCode,
			)
			c.Redirect(http.StatusFound, redirectURL)
			return
		}
 
		err = repo.HandlePaymentResult(orderID, transactionNo, rspCode, amount)
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{"code": "97", "message": "Order processing failed"})
			redirectURL := fmt.Sprintf(
				"%s?amount=%f&date=%s&infor=%s&response-code=%s",
				client_url_fail,
				floatAmount,
				payment_content,
				payment_date,
				rspCode,
			)
			c.Redirect(http.StatusFound, redirectURL)
			return
		}

		redirectURL := fmt.Sprintf(
			"%s?amount=%f&date=%s&infor=%s",
			client_url_success,
			floatAmount,
			payment_content,
			payment_date,
		)
		c.Redirect(http.StatusFound, redirectURL)
	}
}
