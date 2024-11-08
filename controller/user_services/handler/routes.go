package userhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/users")
	{
		eg.POST("/authentication", authenticationHandler(db))
		eg.POST("/sign-up", registerHandler(db))
	}
}
