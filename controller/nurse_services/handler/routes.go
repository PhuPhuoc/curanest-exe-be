package nursehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterNurseRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/nurses")
	{
		eg.GET("", getNursesHandler(db))
		eg.GET("/:nurse_id", getDetailNurseHandler(db))
		eg.POST("", createNewNurseHandler(db))
		eg.POST("/:nurse_id/register-weekly-work-schedule", registerWeeklyScheduleHandler(db))
		eg.GET("/:nurse_id/get-weekly-work-schedule", getSchedules(db))
	}
}
