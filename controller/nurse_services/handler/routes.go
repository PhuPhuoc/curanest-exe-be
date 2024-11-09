package nursehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterNurseRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/nurses")
	{
		eg.GET("", getNursesHandler(db))
		eg.POST("", createNewNurseHandler(db))
	}
}
