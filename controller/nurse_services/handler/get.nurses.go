package nursehandler

import (
	"net/http"

	nurserepository "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// @BasePath		/api/v1
// @Summary		get list nurses (card)
// @Description	get list nurses (card)
// @Tags			nurses
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/nurses [get]
func getNursesHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := nurserepository.NewNurseStore(db)
		list_fate, err := repo.GetNurses()
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list categories", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_fate, nil)
	}
}
