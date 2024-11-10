package nursehandler

import (
	"net/http"

	nurserepository "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		get list nurses (card)
//	@Description	get list nurses (card)
//	@Tags			nurses
//	@Accept			json
//	@Produce		json
//	@Param			nurse_id	path		string					true	"Nurse ID"
//	@Param			role		query		string					false	"Role"	Enums(user, admin)
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/nurses/{nurse_id} [get]
func getDetailNurseHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nurse_id := c.Param("nurse_id")
		if nurse_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
			return
		}
		role := c.DefaultQuery("role", "user") // defaults to "user" if not provided
		repo := nurserepository.NewNurseStore(db)
		nurses, err := repo.GetNurseDetail(nurse_id, role)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "Cannot get list of nurses", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", nurses, nil)
	}
}
