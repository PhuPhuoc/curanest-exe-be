package userhandler

import (
	"net/http"

	usermodel "github.com/PhuPhuoc/curanest_exe_be/controller/user_services/model"
	userrepository "github.com/PhuPhuoc/curanest_exe_be/controller/user_services/repository"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath		/api/v1
//	@Summary		User login sign in
//	@Description	User login sign in
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user-login-data	body		usermodel.LoginForm		true	"User log in info"
//	@Success		200				{object}	map[string]interface{}	"user data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/users/authentication [post]
func authenticationHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.LoginForm
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, http.StatusBadRequest, "invalid request body", "")
			return
		}
		repo := userrepository.NewUserStore(db)
		user, err := repo.Login(&req)
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, "cannot login now", err.Error())
			return
		}
		utils.SendSuccess(c, http.StatusOK, "login sucessfully", user, nil)
	}
}
