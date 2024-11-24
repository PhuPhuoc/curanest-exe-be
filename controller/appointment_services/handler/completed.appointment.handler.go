package appointmenthandler

import (
	"net/http"

	appointmentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		nurse completed appointment
// @Description	nurse completed appointment
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			appointment_id	path		string					true	"Appointment ID"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/appointments/{appointment_id}/completed [post]
func completedAppointment(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appointment_id := c.Param("appointment_id")
		if appointment_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment ID is required"})
			return
		}
		repo := appointmentrepository.NewAppoinmentStore(db)
		err := repo.CompletedAppointment(appointment_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot completed this appointment", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "completed appointment successfully", nil, nil)
	}
}
