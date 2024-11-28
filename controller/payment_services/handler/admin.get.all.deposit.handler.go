package paymenthandler

import (
	"net/http"

	paymentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get all deposit transaction
//	@Description	get all deposit transaction
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"message success"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/payments/admin/deposit-transactions [get]
func adminGetAllDepositTransaction(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := paymentrepository.NewPaymentStore(db)
		data, err := repo.GetAllDepositTransactionOfNurse()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get all deposit transaction", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get data successful", data, nil)
	}
}
