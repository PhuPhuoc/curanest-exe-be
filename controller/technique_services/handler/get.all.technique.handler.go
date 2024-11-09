package techniquehandler

import (
	"net/http"

	techniquerepository "github.com/PhuPhuoc/curanest_exe_be/controller/technique_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get techniques
// @Description	get techniques
// @Tags			techniques
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/techniques [get]
func getTechniquesHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := techniquerepository.NewTechniqueStore(db)
		list_post, err := repo.GetAllTechnique()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list techniques", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_post, nil)
	}
}
