package nursehandler

import (
	"net/http"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
	nurserepository "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		register for nurse's work schedule
//	@Description	register for nurse's work schedule
//	@Tags			nurses
//	@Accept			json
//	@Produce		json
//	@Param			nurse_id		path		string							true	"Nurse ID"
//	@Param			creation_form	body		nursemodel.WeeklyWorkSchedule	true	"Nurse shift information"
//	@Success		200				{object}	map[string]interface{}			"data"
//	@Failure		400				{object}	error							"Bad request error"
//	@Router			/nurses/{nurse_id}/register-weekly-work-schedule [post]
func registerWeeklyScheduleHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nurse_id := c.Param("nurse_id")
		if nurse_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
			return
		}

		var req nursemodel.WeeklyWorkSchedule
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}

		repo := nurserepository.NewNurseStore(db)
		err := repo.RegisterWeeklySchedule(nurse_id, &req)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "Schedule registration failed", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Schedule registration successful", nil, nil)
	}
}
