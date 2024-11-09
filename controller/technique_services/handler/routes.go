package techniquehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterTechniqueRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/techniques")
	{
		eg.GET("", getTechniquesHandler(db))
		eg.POST("", createTechniqueHandler(db))
	}
}
