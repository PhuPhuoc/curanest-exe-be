package customerhandler

import (
	"net/http"

	customerrepository "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get customer information
//	@Description	get customer information
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			customer_id	path		string					true	"Customer ID"
//	@Success		200			{object}	map[string]interface{}	"message success"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/customers/{customer_id} [get]
func getCustomerInfoHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer_id := c.Param("customer_id")
		if customer_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "customer ID is required"})
			return
		}
		repo := customerrepository.NewCustomerStore(db)
		data, err := repo.GetCustomerInfo(customer_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get data successful", data, nil)
	}
}
