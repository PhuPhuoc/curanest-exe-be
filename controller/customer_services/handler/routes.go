package customerhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterCustomerRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	rg.POST("/users/:user_id/create-customer-profile", createCustomerHandler(db))

	eg := rg.Group("/customers")
	{
		eg.GET("/:customer_id", getCustomerInfoHandler(db))
	}
}
