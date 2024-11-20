package nursehandler

import (
	"net/http"

	nurserepository "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

)

//	@BasePath		/api/v1
//	@Summary		get weeky shedule of nurses
//	@Description	get weeky shedule of nurses
//	@Tags			nurses
//	@Accept			json
//	@Produce		json
//	@Param			nurse_id	path		string					true	"Nurse ID"
//	@Param			from		query		string					true	"date from"
//	@Param			to			query		string					true	"date to"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/nurses/{nurse_id}/get-weekly-work-schedule [get]
func getSchedules(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nurse_id := c.Param("nurse_id")
		if nurse_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
			return
		}

		from := c.Query("from")
		to := c.Query("to")

		if from == "" || to == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters 'from', 'to'"})
			return
		}

		repo := nurserepository.NewNurseStore(db)
		list_fate, err := repo.GetSchedules(nurse_id, from, to)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get nurse's schedule", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_fate, nil)
	}
}
