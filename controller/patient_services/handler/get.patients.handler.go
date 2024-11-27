package patienthandler

import (
	"net/http"

	patientrepository "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get list patients of customer
// @Description	get list patients of customer
// @Tags			patients
// @Accept			json
// @Produce		json
// @Param			customer_id	path		string					true	"Customer ID"
// @Success		200			{object}	map[string]interface{}	"message success"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/customers/{customer_id}/patients [get]
func getPatientsHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer_id := c.Param("customer_id")
		if customer_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
			return
		}
		repo := patientrepository.NewPatientStore(db)
		data, err := repo.GetNurseByCustomerID(customer_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get data successfully", data, nil)
	}
}
