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
//	@Summary		get list nurses (card)
//	@Description	get list nurses (card)
//	@Tags			nurses
//	@Accept			json
//	@Produce		json
//	@Param			full-name		query		string					false	"Filter by nurse's full name"
//	@Param			phone-number	query		string					false	"Filter by nurse's phone number"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/nurses [get]
func getNursesHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get query from url
		fullName := c.Query("full_name")
		phoneNumber := c.Query("phone_number")

		filterParams := &nursemodel.NurseFilter{
			FullName:    fullName,
			PhoneNumber: phoneNumber,
		}

		repo := nurserepository.NewNurseStore(db)
		list_fate, err := repo.GetNurses(filterParams)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot get list nurses", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "Get data successfully", list_fate, nil)
	}
}
