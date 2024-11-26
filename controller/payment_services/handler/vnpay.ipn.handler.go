package paymenthandler

import (
	"net/http"

	paymentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

)

func VNPayIPNHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy tất cả các tham số từ query
		vnpParams := c.Request.URL.Query()

		// Gọi repository để xử lý IPN
		repo := paymentrepository.NewPaymentStore(db)
		err := repo.ProcessVNPayIPN(vnpParams)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "Failed to process VNPay IPN", err.Error())
			return
		}

		// Phản hồi thành công
		utils.SendSuccess(c, http.StatusOK, "VNPay IPN processed successfully", gin.H{
			"RspCode": "00",
			"Message": "Success",
		}, nil)
	}
}
