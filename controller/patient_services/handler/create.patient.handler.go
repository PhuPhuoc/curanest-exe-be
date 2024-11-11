package patienthandler

import (
	"net/http"

	patientmodel "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/model"
	patientrepository "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		create new customer
//	@Description	create new customer
//	@Tags			patients
//	@Accept			json
//	@Produce		json
//	@Param			customer_id				path		string					true								"Customer ID"
//	@Param			customer-creation-data	info		body					patientmodel.PatientCreationModel	true	"Patient register data"
//	@Success		200						{object}	map[string]interface{}	"message success"
//	@Failure		400						{object}	error					"Bad request error"
//	@Router			/customers/{customer_id}/create-patient-profile [post]
func createPatientHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer_id := c.Param("customer_id")
		if customer_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
			return
		}
		var req patientmodel.PatientCreationModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := patientrepository.NewPatientStore(db)
		if err := repo.CreatePatient(customer_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new market's post", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new patient registration successful", nil, nil)
	}
}
