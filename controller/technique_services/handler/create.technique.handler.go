package techniquehandler

import (
	"net/http"

	techniquemodel "github.com/PhuPhuoc/curanest_exe_be/controller/technique_services/model"
	techniquerepository "github.com/PhuPhuoc/curanest_exe_be/controller/technique_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		create new technique
//	@Description	create new technique
//	@Tags			techniques
//	@Accept			json
//	@Produce		json
//	@Param			technique_form	body		techniquemodel.TechniqueCreationModel	true	"new technique information"
//	@Success		200				{object}	map[string]interface{}					"message success"
//	@Failure		400				{object}	error									"Bad request error"
//	@Router			/techniques [post]
func createTechniqueHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req techniquemodel.TechniqueCreationModel
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := techniquerepository.NewTechniqueStore(db)
		if err := repo.CreateTechnique(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot create new technique", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "new technique registration successful", nil, nil)
	}
}
