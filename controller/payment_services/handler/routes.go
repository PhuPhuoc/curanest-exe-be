package paymenthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPaymentRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("payments/")
	{
		eg.POST("", createPayment(db))
		eg.GET("/vnpay_ipn", VNPayIPNHandler(db))
		eg.GET("/vnpay_return", VNPayReturnHandler(db))
		eg.GET("/vnpay-url", createPaymentVer2(db))
		eg.GET("/current-wallet-amount/:user_id", getWalletAmount(db))
		eg.GET(":user_id/deposit-transactions", getAllDepositTransaction(db))
		eg.GET(":user_id/wallet-transactions", getAllWalletTransaction(db))
		eg.GET("admin/deposit-transactions", adminGetAllDepositTransaction(db))
		eg.GET("admin/wallet-transactions", adminGetAllWalletTransaction(db))
	}
}
