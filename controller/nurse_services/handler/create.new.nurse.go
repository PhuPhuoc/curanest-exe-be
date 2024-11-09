package nursehandler

import (
	"net/http"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
	nurserepository "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		create new nurse
// @Description	create new nurse
// @Tags			nurses
// @Accept			json
// @Produce		json
// @Param			creation_form	body		nursemodel.NurseCreationModel	true	"Nurse creation data"
// @Success		200				{object}	map[string]interface{}			"message success"
// @Failure		400				{object}	error							"Bad request error"
// @Router			/nurses [post]
func createNewNurseHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req nursemodel.NurseCreationModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := nurserepository.NewNurseStore(db)
		if err := repo.CreateNurse(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new nurse", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new nurse registration successful", nil, nil)
	}
}
