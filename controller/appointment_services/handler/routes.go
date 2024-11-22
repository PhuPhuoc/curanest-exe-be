package appointmenthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

)

func RegisterNurseRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("appointments/")
	{
		eg.POST("", registerAppointment(db))
		eg.POST("/:appointment_id/confirm", confirmAppointment(db))
		eg.GET("/:appointment_id", getAppointmentDetail(db))
	}
}