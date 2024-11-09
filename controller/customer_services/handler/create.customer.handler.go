package customerhandler

import (
	"net/http"

	customermodel "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/model"
	customerrepository "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		create new customer
// @Description	create new customer
// @Tags			customer
// @Accept			json
// @Produce		json
// @Param			user_id					path		string					true						"User ID"
// @Param			customer-creation-data	info		body					customermodel.CustomerModel	true	"Customer register data"
// @Success		200						{object}	map[string]interface{}	"message success"
// @Failure		400						{object}	error					"Bad request error"
// @Router			/users/{user_id}/create-customer-profile [post]
func createCustomerHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		var req customermodel.CustomerModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := customerrepository.NewCustomerStore(db)
		if err := repo.CreateCustomer(user_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new account registration successful", nil, nil)
	}
}
