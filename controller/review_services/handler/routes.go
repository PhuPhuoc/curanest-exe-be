package reviewhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterReviewRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	rg.POST("appointments/:appointment_id/reviews", createNewReview(db))
	rg.GET("appointments/nurse/:nurse_id/reviews", getAllReview(db))
	// /appointments/nurse/{nurse_id}/reviews
}
