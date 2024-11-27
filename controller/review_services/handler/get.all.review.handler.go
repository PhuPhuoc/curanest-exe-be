package reviewhandler

import (
	"net/http"

	reviewrepository "github.com/PhuPhuoc/curanest_exe_be/controller/review_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get all review
// @Description	get all review
// @Tags			reviews
// @Accept			json
// @Produce		json
// @Param			nurse_id	path		string					true	"nurse id"
// @Success		200			{object}	map[string]interface{}	"message success"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/appointments/nurse/{nurse_id}/reviews [get]
func getAllReview(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nurse_id := c.Param("nurse_id")
		if nurse_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nurse ID is required"})
			return
		}
		repo := reviewrepository.NewReviewStore(db)
		data, err := repo.GetAllReviewOrNurse(nurse_id)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new review", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "get data successful", data, nil)
	}
}
