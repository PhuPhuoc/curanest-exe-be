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

	}
}
