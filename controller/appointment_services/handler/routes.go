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
		eg.POST("/:appointment_id/completed", completedAppointment(db))
		eg.GET("/:appointment_id", getAppointmentDetail(db))
		eg.GET("/card/:account_id", getAppointmentCard(db))
	}
}
