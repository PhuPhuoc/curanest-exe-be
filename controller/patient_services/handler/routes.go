package patienthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPatientRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	rg.POST("/customers/:customer_id/create-patient-profile", createPatientHandler(db))
	// eg := rg.Group("/patients")
	// {

	// }
}
