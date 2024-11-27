package paymenthandler

import (
	"net/http"

	paymentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get current amount in wallet of nurse
// @Description	get current amount in wallet of nurse
// @Tags			payments
// @Accept			json
// @Produce		json
// @Param			user_id	path		string					true	"user id"
// @Success		200		{object}	map[string]interface{}	"message success"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/payments/current-wallet-amount/{user_id} [get]
func getWalletAmount(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := paymentrepository.NewPaymentStore(db)
		data, err := repo.GetWalletAmount(user_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get current amount in wallet", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get data successful", data, nil)
	}
}
