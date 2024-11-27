package reviewhandler

import (
	"net/http"

	reviewmodel "github.com/PhuPhuoc/curanest_exe_be/controller/review_services/model"
	reviewrepository "github.com/PhuPhuoc/curanest_exe_be/controller/review_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		create new review
// @Description	create new review
// @Tags			reviews
// @Accept			json
// @Produce		json
// @Param			appointment_id	path		string								true	"appointment id"
// @Param			create_review	body		reviewmodel.CreateNewReviewModel	true	"create new review in appointment"
// @Success		200				{object}	map[string]interface{}				"message success"
// @Failure		400				{object}	error								"Bad request error"
// @Router			/appointments/{appointment_id}/reviews [post]
func createNewReview(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req reviewmodel.CreateNewReviewModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		appointment_id := c.Param("appointment_id")
		if appointment_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment ID is required"})
			return
		}
		repo := reviewrepository.NewReviewStore(db)
		if err := repo.CreateNewReview(appointment_id, &req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new review", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new review created successful", nil, nil)
	}
}
