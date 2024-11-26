package paymenthandler

import (
	"log"
	"net/http"
	"os"
	"strconv"

	paymentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/model"
	paymentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

//	@BasePath		/api/v1
//	@Summary		create payment URL
//	@Description	create payment URL
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			amount		query		string					true	"amount"
//	@Param			nurse_id	query		string					true	"nurse_id"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/payments/vnpay-url [get]
func createPaymentVer2(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			log.Println("(dev) Error loading .env file")
			// utils.SendError(c, http.StatusBadRequest, "Error loading .env file", err.Error())
			// return
		}
		vnp_config := &paymentmodel.VnpayConfig{
			TMNCode:   os.Getenv("VNPAY_TMN_CODE"),
			SecretKey: os.Getenv("VNPAY_SECRET_KEY"),
			VNPURL:    os.Getenv("VNPAY_VNPURL"),
			ReturnURL: os.Getenv("VNPAY_RETURN_URL"),
		}
		amount := c.Query("amount")
		nurse_id := c.Query("nurse_id")

		num_amount, _ := strconv.Atoi(amount)
		req := &paymentmodel.CreatePaymentRequest{
			Amount: num_amount,
			UserID: nurse_id,
		}
		// Validate required fields
		if req.Amount <= 0 {
			utils.SendError(c, http.StatusBadRequest, "Amount are required", "Amount are required")
			return
		}

		// Generate VNPAY parameters
		repo := paymentrepository.NewPaymentStore(db)
		paymentURL, err := repo.GenerateVNPayURL(req, vnp_config)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "Failed to generate payment URL", err.Error())
			return
		}

		// Response success
		utils.SendSuccess(c, http.StatusOK, "Payment URL created successfully", gin.H{"payment_url": paymentURL}, nil)
	}
}
