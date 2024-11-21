package appointmenthandler

import (
	"net/http"

	appointmentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/model"
	appointmentrepository "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		create new appointment
// @Description	create new appointment
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			creation_form	body		appointmentmodel.RegisterAppoinment	true	"Nurse creation data"
// @Success		200				{object}	map[string]interface{}				"message success"
// @Failure		400				{object}	error								"Bad request error"
// @Router			/appointments  [post]
func registerAppointment(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req appointmentmodel.RegisterAppoinment
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := appointmentrepository.NewAppoinmentStore(db)
		if err := repo.RegisterAppoinment(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new appointment", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new appointment successful", nil, nil)
	}
}
