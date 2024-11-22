package appointmenthandler

import (
	"net/http"

	appointmentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		nurse confirm appointment
//	@Description	nurse confirm appointment
//	@Tags			appointments
//	@Accept			json
//	@Produce		json
//	@Param			appointment_id	path		string					true	"Appointment ID"
//	@Param			confirm			query		string					true	"confirm"	Enums(confirmed, cancel)
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/appointments/{appointment_id}/confirm [post]
func confirmAppointment(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appointment_id := c.Param("appointment_id")
		if appointment_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment ID is required"})
			return
		}
		status := c.Query("confirm")
		if status != "confirmed" && status != "cancel" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "confirm query must be 'confirmed' or 'cancel'"})
			return
		}
		repo := appointmentrepository.NewAppoinmentStore(db)
		err := repo.ConfirmAppointment(appointment_id, status)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot confirm this appointment", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "confirm appointment successfully", nil, nil)
	}
}
