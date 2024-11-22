package appointmenthandler

import (
	"net/http"

	appointmentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get appointment (card)
//	@Description	get appointment (card)
//	@Tags			appointments
//	@Accept			json
//	@Produce		json
//	@Param			account_id	path		string					true	"Account ID"
//	@Param			role		query		string					true	"Role"	Enums(patient, nurse)
//	@Param			from		query		string					true	"date from"
//	@Param			to			query		string					true	"date to"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/appointments/card/{account_id} [get]
func getAppointmentCard(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		account_id := c.Param("account_id")
		if account_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
			return
		}
		role := c.Query("role")
		if role != "patient" && role != "nurse" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role query must be 'patient' or 'nurse'"})
			return
		}
		from := c.Query("from")
		to := c.Query("to")
		repo := appointmentrepository.NewAppoinmentStore(db)
		data, err := repo.GetAppointmentCard(account_id, role, from, to)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "Cannot get nurse's data", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", data, nil)
	}
}
